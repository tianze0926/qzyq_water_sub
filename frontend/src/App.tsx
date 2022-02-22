import {
  FormEvent, useCallback, useEffect, useState,
} from 'react';
import {
  Box, Button, Grid, TextField, AlertColor,
} from '@mui/material';
import { DataGrid, GridColDef, GridValueFormatterParams } from '@mui/x-data-grid';

import axios, { AxiosError, AxiosResponse } from 'axios';

import Notice, { NoticeProps } from './Notice';

const recordApi = axios.create({
  baseURL: 'api/record',
});

type Record = {
  id: number;
  time: number;
  ip: string;
  file_number: string;
}

const columns: GridColDef[] = [
  {
    field: 'file_number',
    headerName: '档案号',
    flex: 3,
    minWidth: 50,
  },
  {
    field: 'time',
    headerName: '时间',
    valueFormatter: (params: GridValueFormatterParams) => (
      new Date(params.value as number).toLocaleString()
    ),
    flex: 4,
    minWidth: 200,
  },
  {
    field: 'ip',
    headerName: 'IP',
    flex: 4,
    minWidth: 100,
  },
].map((c) => ({ ...c, sortable: false }));

type RowsState = {
  page: number;
  pageSize: number;
  rows: Record[];
  loading: boolean;
  rowCount: number;
}

function App() {
  const [fileNumber, setFileNumber] = useState('11213');
  const [submitTimes, setSubmitTimes] = useState(0); // for refreshing the table

  const [rowsState, setRowsState] = useState<RowsState>({
    page: 0,
    pageSize: 5,
    rows: [],
    loading: false,
    rowCount: 0,
  });

  const [notice, setNotice] = useState<NoticeProps>({
    timestamp: new Date().getTime(),
    message: '',
    severity: undefined,
  });
  const notify = useCallback((message: string, severity?: AlertColor) => (
    setNotice({
      timestamp: new Date().getTime(),
      message,
      severity,
    })
  ), []);

  const handleError = useCallback((error: AxiosError) => {
    if (error.response) {
      notify(`Status: ${error.response.status} Data: ${error.response.data}`, 'error');
    } else if (error.request) {
      notify('No response was received', 'error');
    } else {
      notify(`Something happened in setting up the request: ${error.message}`, 'error');
    }
  }, [notify]);

  const handleSubmit = (e: FormEvent) => {
    recordApi.post('', { file_number: fileNumber })
      .then(() => {
        notify(`提交成功，档案号：${fileNumber}`, 'success');
        setSubmitTimes((prev) => prev + 1);
      })
      .catch(handleError);
    e.preventDefault();
  };

  useEffect(() => {
    let active = true;
    setRowsState((prev) => ({ ...prev, loading: true }));
    recordApi.get('', {
      params: {
        page: rowsState.page,
        size: rowsState.pageSize,
      },
    }).then((response: AxiosResponse<{data: Record[], total: number}>) => {
      if (!active) return;
      setRowsState((prev) => ({
        ...prev,
        loading: false,
        rows: response.data.data,
        rowCount: response.data.total,
      }));
    })
      .catch((error: AxiosError) => {
        setRowsState((prev) => ({ ...prev, loading: false }));
        handleError(error);
      });
    return () => {
      active = false;
    };
  }, [rowsState.page, rowsState.pageSize, submitTimes, handleError]);

  return (
    <div className="App">
      <Box my={2} mx={4}>
        <Grid
          container
          spacing={{ xs: 2, md: 3 }}
          columns={{ xs: 4, sm: 4, md: 12 }}
          justifyContent="center"
        >
          <Grid item xs={4} sm={4} md={4}>
            <Box component="form" onSubmit={handleSubmit} sx={{ maxWidth: '600px', mx: 'auto' }}>
              <TextField
                fullWidth
                variant="filled"
                label="档案号"
                value={fileNumber}
                onChange={(event) => setFileNumber(event.target.value)}
              />
              <Box mt={1}>
                <Button type="submit" variant="contained">提交</Button>
              </Box>
            </Box>

          </Grid>
          <Grid item xs={4} sm={4} md={6}>
            <DataGrid
              sx={{ minWidth: '450px', maxWidth: '600px', mx: 'auto' }}
              columns={columns}
              autoHeight
              disableSelectionOnClick
              disableColumnMenu
              rowsPerPageOptions={[5, 10, 20]}
              page={rowsState.page}
              pageSize={rowsState.pageSize}
              rows={rowsState.rows}
              loading={rowsState.loading}
              rowCount={rowsState.rowCount}
              paginationMode="server"
              onPageChange={(page) => setRowsState((prev) => ({ ...prev, page }))}
              onPageSizeChange={(pageSize) => setRowsState((prev) => ({ ...prev, pageSize }))}
            />
          </Grid>
        </Grid>
      </Box>
      <Notice
        timestamp={notice.timestamp}
        message={notice.message}
        severity={notice.severity}
      />
    </div>
  );
}

export default App;
