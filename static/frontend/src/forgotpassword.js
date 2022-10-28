import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
// import FormControlLabel from '@mui/material/FormControlLabel';
// import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';

// import { Trans } from 'react-i18next';

const theme = createTheme();

const ForgotPassword = (props) => {
 
    const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    console.log({
      email: data.get('login'),
    
    });
    
  };

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
          Forgot password?
          </Typography>
          <Typography component="body" sx={{mt: 3}}>
          Enter the email address associated with your account and we will send you a link to reset your password.
          </Typography>
         

          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 4 }}>
                <TextField
                  required
                  fullWidth
                  id="login"
                  label="Email Address or Mobile"
                  name="login"
                  autoComplete="login"
                />

            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Request password reset
            </Button>


            <Grid container justifyContent="flex-start">
              <Grid item>
                <Link href="/signin" variant="body2">
                  I remember password ...
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}


export {ForgotPassword}