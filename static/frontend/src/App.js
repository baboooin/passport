import React, {useState} from 'react';
import { styled, createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import MuiDrawer from '@mui/material/Drawer';
import Box from '@mui/material/Box';
import MuiAppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import Badge from '@mui/material/Badge';
import Container from '@mui/material/Container';
// import Grid from '@mui/material/Grid';
// import Paper from '@mui/material/Paper';
// import Link from '@mui/material/Link';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import NotificationsIcon from '@mui/icons-material/Notifications';
// import { mainListItems, SecondaryListItems } from './DrawerMenu';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import DashboardIcon from '@mui/icons-material/Dashboard';
import UserIcon from '@mui/icons-material/Person';
import {isMobile} from 'react-device-detect';

import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
// import PeopleIcon from '@mui/icons-material/People';
// import BarChartIcon from '@mui/icons-material/BarChart';
// import LayersIcon from '@mui/icons-material/Layers';
// import AssignmentIcon from '@mui/icons-material/Assignment';
import {Link} from 'react-router-dom';
import {TasksIcon, CompanyIcon, WarehouseIcon, UsersIcon} from './SvgIcons';

// import useToken from './useToken';

import {Routes, Route, useLocation} from 'react-router-dom';

import Home from './home';
import About from './about';
import {SignUp} from './signUp';
import {SignIn} from './signIn';
import {ForgotPassword} from './forgotpassword';

import {PassportIcon} from './SvgIcons';


const drawerWidth = 240;

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(['width', 'margin'], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const Drawer = styled(MuiDrawer, { shouldForwardProp: (prop) => prop !== 'open' })(
  ({ theme, open }) => ({
    '& .MuiDrawer-paper': {
      position: 'relative',
      whiteSpace: 'nowrap',
      width: drawerWidth,
      transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.enteringScreen,
      }),
      boxSizing: 'border-box',
      ...(!open && {
        overflowX: 'hidden',
        transition: theme.transitions.create('width', {
          easing: theme.transitions.easing.sharp,
          duration: theme.transitions.duration.leavingScreen,
        }),
        width: theme.spacing(7),
        [theme.breakpoints.up('sm')]: {
          width: theme.spacing(9),
        },
      }),
    },
  }),
);

const mdTheme = createTheme();




const Main  =  (props) => {
  
  const [open, setOpen] = React.useState(!isMobile);
  const toggleDrawer = () => {
    setOpen(!open);
  };
return(
  <Box sx={{ display: 'flex' }}>
  <CssBaseline />
  <AppBar position="absolute" open={open}>
    <Toolbar
      sx={{
        pr: '24px', // keep right padding when drawer closed
      }}
    >
      <IconButton
        edge="start"
        color="inherit"
        aria-label="open drawer"
        onClick={toggleDrawer}
        sx={{
          marginRight: '36px',
          ...(open && { display: 'none' }),
        }}
      >
        <MenuIcon />
      </IconButton>
      <PassportIcon color = "white" sx={{ display: { xs: 'none', md: 'flex' }, mr: 2 }}/>
    <Typography variant="h5" noWrap component="div"
     sx={{
      mr: 2,
      display: { xs: 'flex', md: 'flex' },
      fontFamily: 'monospace',
      fontWeight: 700,
      letterSpacing: '.2rem',
      color: 'inherit',
      textDecoration: 'none',
    }}
    >
      PASSPORT
    </Typography>

    <Box sx={{ flexGrow: 1 }} />
      <IconButton color="inherit">
        <Badge badgeContent={4} color="secondary">
          <NotificationsIcon />
        </Badge>
      </IconButton>
    </Toolbar>
  </AppBar>

  <Drawer variant="permanent" open={open}>
    <Toolbar
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
        px: [1],
      }}
    >
      <IconButton onClick={toggleDrawer}>
        <ChevronLeftIcon />
      </IconButton>
    </Toolbar>
    <Divider />
    <List component="nav">
    <React.Fragment>
    <ListItemButton component={Link} to='/'>
      <ListItemIcon>
        <DashboardIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Dashboard" />
    </ListItemButton>


    <ListItemButton component={Link} to='/about'>
      <ListItemIcon>
        <TasksIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Task" />
    </ListItemButton>

    <ListItemButton>
      <ListItemIcon>
        <ShoppingCartIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Orders" />
    </ListItemButton>
    </React.Fragment>
      <Divider sx={{ my: 1 }} />
      <React.Fragment>
    <ListSubheader component="div" inset>
      Controls
    </ListSubheader>
    <ListItemButton component = {Link} to="SignUp">
      <ListItemIcon>
        <UserIcon color = "#04292A" />
      </ListItemIcon>
      <ListItemText primary="Profile" />
    </ListItemButton>

    <ListItemButton component = {Link} to="SignUp">
      <ListItemIcon>
        <CompanyIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Companies" />
    </ListItemButton>

    <ListItemButton component = {Link} to="SignUp">
      <ListItemIcon>
        <WarehouseIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Warehouse" />
    </ListItemButton>
 
 
    <ListItemButton component = {Link} to="SignUp">
      <ListItemIcon>
        <UsersIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Users" />
    </ListItemButton>
 
    <ListItemButton onClick={() => {
        props.setToken(null)
        localStorage.removeItem('token')
    }}>
      <ListItemIcon>
        <UsersIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Exit" />
    </ListItemButton>
 

  </React.Fragment>
    </List>
  </Drawer>

  <Box
    component="main"
    sx={{
      backgroundColor: (theme) =>
        theme.palette.mode === 'light'
          ? theme.palette.grey[100]
          : theme.palette.grey[900],
      flexGrow: 1,
      height: '100vh',
      overflow: 'auto',
    }}
  >
    <Toolbar />
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="about" element={<About />} />
      <Route path="signup" element={<SignUp />} />
    </Routes> 
    </Container>  
   
  </Box>
</Box>

);
};




export default function App() {

  // const authContext = createContext();

  // const {token, setToken} = useToken();  

  const [token, setToken] = useState(localStorage.getItem("token"));

  console.log("token = ", token);

  
  let location = useLocation();

if (location.pathname === '/ForgotPassword') {
  return <ForgotPassword />
}

if (location.pathname === '/SignUp') {
  return <SignUp />
}

  if(!token) {
    return <SignIn setToken={setToken} />
  }

  return (
    <ThemeProvider theme={mdTheme}>
    <Main setToken = {setToken}/>
    </ThemeProvider>
  );
}




