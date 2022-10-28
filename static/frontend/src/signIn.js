import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
// import FD2JSON from './FormDataToJson';
import LocalizedStrings from 'react-localization';


let t = new LocalizedStrings({
 en:{
  SignIn:"Sign In",
  loginlabel:"Email Address or Mobile",
  passwordlabel:"Password",
  ForgotPassword:"Forgot password",
  RememberMe:"Remember me",
 },
 ru: {
  SignIn:"Авторизация",
  loginlabel:"Электрическая почта или номер сотового",
  passwordlabel:"Пароль",
  ForgotPassword:"Просклерозил",
  RememberMe:"Запомнить",
  
 },
 uk: {
  SignIn:"Авторизацiя",
  loginlabel:"Електронна пошта або мобільний телефон",
  passwordlabel:"Пароль",
  ForgotPassword:"Забув",
  RememberMe:"Запам`ятати",
 },
});

// import { Trans } from 'react-i18next';

const theme = createTheme();

const SignIn = (props) => {
 
    const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    console.log(JSON.stringify(Object.fromEntries(data)));
    const token = data.get('login'); // temp 
    localStorage.setItem('token',token);
    props.setToken(token);
    
    console.log({
      login: data.get('login'),
    });
    
  };
  let l = navigator.language.substring(0,2) || navigator.userLanguage.substring(0,2)
  
  console.log(l);
  t.setLanguage(l)
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
            {t.SignIn}
          </Typography>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
            <Grid container spacing={2}>
              
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="login"
                  label={t.loginlabel} //"Email Address or Mobile"
                  name="login"
                  autoComplete="login"
                />
              </Grid>
            
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  name="password"
                  label={t.passwordlabel}
                  type="password"
                  id="password"
                  autoComplete="password"
                />
              </Grid>

              <Grid item xs={12}>
                <FormControlLabel
                  control={<Checkbox value="remember" color="primary" />}
                  label={t.RememberMe}
                />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              {t.SignIn}
            </Button>

            <Grid container justifyContent="flex-start">
              <Grid item>
                <Link href="/ForgotPassword" variant="body2">
                {t.ForgotPassword}?
                </Link>
              </Grid>
            </Grid>

            <Grid container justifyContent="flex-start">
              <Grid item>
                <Link href="/signup" variant="body2">
                  No Account ? Sign up
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}


export {SignIn}