import * as React from "react";
import { Link } from "react-router-dom";
import Toolbar from '@mui/material/Toolbar';
// import Typography from '@mui/material/Typography';
// import {PassportIcon} from './SvgIcons'
import Button from '@mui/material/Button';
// import Container from '@mui/material/Container';
import Box from '@mui/material/Box';

export default function Home() {


  const pages = [
    {title:"API",
     link: "/api"
    },
    {title:"Sign In",
    link: "/signIn"
   },
   {title:"Sign Up",
   link: "/signUp"
  },
    ];
    
    return (

      
      <React.Fragment>
<Toolbar sx={{background:"red"}}>
             <Box sx={{ flexGrow: 1 }} />
          <Box sx={{display: { xs: 'none', md: 'flex' } , mr: 2 }}>
             {pages.map((page) => (
               <Button
                component = {Link}
                 to = {page.link}
                 key={page.title}
                 sx={{ my: 0, color: 'white', display: 'block' }}
               >
                 {page.title}
               </Button>
             ))}
           </Box>
         </Toolbar>
        <main>
          <h2>Welcome to the homepage!</h2>
          <p>You can do this, I believe in you.</p>
        </main>
        <nav>
          <Link to="/about">About</Link>
        </nav>
      </React.Fragment>
    );
  }

  