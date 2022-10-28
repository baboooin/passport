
import Divider from '@mui/material/Divider';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
// import InboxIcon from '@mui/icons-material/MoveToInbox';
import MailIcon from '@mui/icons-material/Mail';
// import HomeIcon from '@mui/icons-material/Home';
import {CompanyIcon} from './SvgIcons'

import List from '@mui/material/List';
import Drawer from '@mui/material/Drawer';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';

import { Link } from "react-router-dom";
import Typography from '@mui/material/Typography';
import {PassportIcon} from './SvgIcons'
// import Button from '@mui/material/Button';


const drawerWidth = 240;


const MyDrawer = (props) =>{
  return (
    <Drawer
    variant="persistent"
    open = {props.open}

    sx={{
      width: drawerWidth,
      flexShrink: 0,
      border:"none",
      background:"#04292A",
      [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
    }}
  >
    <Toolbar background = {props.background}>
          <PassportIcon color = "black" sx={{ display: { xs: 'none', md: 'flex' }, mr: 2 }}/>
          <Typography variant="h5" noWrap component="div"
           sx={{
            mr: 2,
            display: { xs: 'none', md: 'flex' },
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
          
      </Toolbar>

    <Box sx={{ overflow: 'auto' }}>
      <List>
        <ListItem key="Home" button disablePadding  component={Link} to="/">
        <ListItemButton>
        <ListItemIcon>
        <CompanyIcon />
            </ListItemIcon>
            <ListItemText primary="Home"/>
            </ListItemButton>
     
        </ListItem>
      </List>
      <Divider />
      <List>
    
          <ListItem key="About" disablePadding component={Link} to="about">
            <ListItemButton>
              <ListItemIcon>
                <MailIcon/>
              </ListItemIcon>
              <ListItemText primary="About" />
            </ListItemButton>
          </ListItem>
    
      </List> 
    </Box>
  </Drawer>
);

}

export default MyDrawer