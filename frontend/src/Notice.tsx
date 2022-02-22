import { useState, useEffect, useRef } from 'react';
import Snackbar, { SnackbarOrigin } from '@mui/material/Snackbar';
import { Alert, AlertColor } from '@mui/material';

const anchorOrigin: SnackbarOrigin = {
  vertical: 'bottom',
  horizontal: 'center',
};

export type NoticeProps = {
  timestamp: number;
  message: string;
  severity: AlertColor | undefined;
}

export default function Notice({ timestamp, message, severity }: NoticeProps) {
  const [open, setOpen] = useState(false);

  // only on updates except initial mount
  const isInitialMount = useRef(true);
  useEffect(() => {
    if (isInitialMount.current) {
      isInitialMount.current = false;
      return () => { };
    }
    setOpen(true);
    return () => setOpen(false);
  }, [timestamp]);

  const handleClose = () => setOpen(false);

  return (severity === undefined)
    ? (
      <Snackbar
        key={timestamp}
        open={open}
        autoHideDuration={2000}
        onClose={handleClose}
        message={message}
        anchorOrigin={anchorOrigin}
      />
    )
    : (
      <Snackbar
        key={timestamp}
        open={open}
        autoHideDuration={2000}
        onClose={handleClose}
        anchorOrigin={anchorOrigin}
      >
        <Alert
            // onClose={handleClose}
          severity={severity}
        >
          {message}
        </Alert>
      </Snackbar>
    );
}
