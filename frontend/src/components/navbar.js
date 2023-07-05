import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { AppBar, Toolbar, IconButton, InputBase, Button } from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';
import '../css/home.css'
import  Cookies from 'universal-cookie';
import { red, blue, grey } from '@mui/material/colors';


const Navbar = () => {
  const cookies = new Cookies();
  const userToken = cookies.get('user');
  const navigate = useNavigate();
  const [searchTerm, setSearchTerm] = useState('');
  const handleLogout = () => {
    cookies.remove('user'); // Eliminar la cookie 'user'
    window.location.replace('/'); // Redirigir al home
  };
  

  const handleSearch = () => {
    if (searchTerm === '') {
      navigate('/');
    } else {
      const formattedSearchTerm = searchTerm.replace(' ', '%20');
      navigate(`/search/${formattedSearchTerm}`);
    }
  };

  const handleInputChange = (e) => {
    setSearchTerm(e.target.value);
  };

  return (

    <AppBar position="static" color="primary">
      <Toolbar>
        <IconButton color="inherit" onClick={handleSearch}>
          <SearchIcon />
        </IconButton>
        <InputBase
          placeholder="Search..."
          value={searchTerm}
          onChange={handleInputChange}
        />
        {userToken ? (
          // Elementos que se mostrarán si el usuario está logueado
          <div>
            <Button onClick={handleLogout} sx={{color: blue[500], transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': {color: red[500],transform: 'scale(1.2)'},position: 'absolute', top: '1',right: '0',justifyContent: 'flex-end', alignItems: 'flex-start', }}>
              LOGOUT
            </Button>
            <Link to="/create">
            <Button  sx={{color: grey[300], transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': {color: blue[500],transform: 'scale(1.1)'}}}>
              CREAR PUBLICACIÓN
            </Button>
            </Link>
            <Link to="/mypublics">
            <Button  sx={{color: grey[300], transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': {color: blue[500],transform: 'scale(1.1)'}}}>
              MIS PUBLICACIONES
            </Button>
            </Link>
          </div>
        ) : (
          // Elementos que se mostrarán si el usuario no está logueado
          <Link to="/login">
            <Button color="inherit">
              LOGIN
            </Button>
          </Link>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;