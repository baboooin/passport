import * as React from 'react';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import DashboardIcon from '@mui/icons-material/Dashboard';
import UserIcon from '@mui/icons-material/Person';

import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
// import PeopleIcon from '@mui/icons-material/People';
// import BarChartIcon from '@mui/icons-material/BarChart';
// import LayersIcon from '@mui/icons-material/Layers';
// import AssignmentIcon from '@mui/icons-material/Assignment';
import {Link} from 'react-router-dom';
import {TasksIcon, CompanyIcon, WarehouseIcon, UsersIcon} from './SvgIcons';


export const mainListItems = (
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


    {/* <ListItemButton>
      <ListItemIcon>
        <PeopleIcon />
      </ListItemIcon>
      <ListItemText primary="Customers" />
    </ListItemButton>
    <ListItemButton>
      <ListItemIcon>
        <BarChartIcon />
      </ListItemIcon>
      <ListItemText primary="Reports" />
    </ListItemButton>
    <ListItemButton>
      <ListItemIcon>
        <LayersIcon />
      </ListItemIcon>
      <ListItemText primary="Integrations" />
    </ListItemButton> */}
  </React.Fragment>
);





export const SecondaryListItems =(

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
 
    <ListItemButton>
      <ListItemIcon>
        <UsersIcon color = "#04292A"/>
      </ListItemIcon>
      <ListItemText primary="Exit" />
    </ListItemButton>
 

  </React.Fragment>
);