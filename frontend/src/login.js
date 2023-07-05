import React, { useState } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import Cookies from "universal-cookie";
import { Link } from 'react-router-dom';
import swal from "sweetalert2";
const Cookie = new Cookies();

const Login = () => {
  const[user,setUser]= useState("");
  const[password,setPassword] = useState("");

  const onChangeUser =  (user)=>{
      setUser(user.target.value);
      
  }
  
  const onChangePas = (password)=>{
  setPassword(password.target.value)};

  
  const requestOptions={
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      
      body: JSON.stringify({username : user, password : password })
  };

  const login = async()=>{
      fetch('http://localhost:9000/user',requestOptions)
      .then(response => {if (response.status == 403) {
         swal.fire({
          text: "Datos incorrectos",
          icon: 'error',
         }).then((result) => {
          if (result.isConfirmed) {
              window.location.reload();
              return response.json()
          }})
      }
      if(response.status==201){
        swal.fire({icon: 'success'}
        ).then((result) => {
          if (result.isConfirmed) {
            window.location.replace("/")
            return response.json()
          }})
      }
      return response.json()})
      .then(response => {
          Cookie.set("user", response.token, {path: "/"})
  })
 
  };
 
  const handleSubmit= (event)=>{
      event.preventDefault();
      login();

  };

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          marginTop: 4,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Link to="/" style={{ textDecoration: 'none' }}>
          <img
            src="https://cdn.freebiesupply.com/logos/large/2x/jl-logo-black-and-white.png"
            alt="Logo"
            style={{ width: 75, cursor: 'pointer', marginTop: 20 }}
          />
        </Link>
        <Typography component="h2" variant="h5" >
          LOGIN 
        </Typography>
        <Box component="form" sx={{ mt: 1 }}>
          <TextField
            label="Username"
            type="text"
            value={user}
            onChange={onChangeUser}
            variant="outlined"
            margin="normal"
            required
            fullWidth
          />
          <TextField
            label="Password"
            type="password"
            value={password}
            onChange={onChangePas}
            variant="outlined"
            margin="normal"
            required
            fullWidth
          />
          <Button
            type="button"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
            onClick={handleSubmit}
          >
            Login
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default Login;
