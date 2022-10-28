
import { Link } from "react-router-dom";
import *  as React from 'react';
// import { Toolbar, Box, Button } from "@mui/material";



export default function About() {
    return (
      <React.Fragment>
        <main>
          <h2>Who are we?</h2>
          <p>
            That feels like an existential question, don't you
            think?
          </p>
        </main>
        <nav>
          <Link to="/">Home</Link>
        </nav>
      </React.Fragment>
    );
  }