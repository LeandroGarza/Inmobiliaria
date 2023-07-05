import React, { useState } from "react";
import './css/App.css';
import { BrowserRouter as Router, Route, Routes, Redirect} from 'react-router-dom';
import Home from './home.js';
import Login from './login.js';
import ProductDetail from './ProductDetails'
import Search from './Search.js';
import Create from './create.js'
import Publics from "./mypublics";


function App() {
  return (
    
    <Router>
      <Routes>
      <Route path="/" element={<Home/>} />
      <Route path="/search/:term" element={<Search />} />
      <Route path="/product/:id" element={<ProductDetail/>} />
      <Route path="/login" element={<Login/>} />
      <Route path="/create" element={<Create/>} />
      <Route path="/mypublics" element={<Publics/>} />
      </Routes>
    </Router>
  );
}

export default App;
