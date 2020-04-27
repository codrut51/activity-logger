import React, { useState } from 'react';
import './nav-bar.css';
import MenuIcon from '@material-ui/icons/Menu'
import { Button } from '@material-ui/core';

function NavBar(props) {

  const [open, setOpen] = useState(false);

  return (
    <div className="NavBar">
        {open && <div className="NavBar-Menu">
                Hello!
        </div>}
      <Button className="NavBar-Button" onClick={() => setOpen(!open)}>
        <MenuIcon className="NavBar-Icon"/>
      </Button>
      <div id="NavBar-Title">

      </div>
    </div>
  );
}

export default NavBar;
