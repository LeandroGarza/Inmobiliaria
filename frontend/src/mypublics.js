import React, { useEffect, useState } from 'react';
import Cookies from "universal-cookie";
const Cookie = new Cookies();

const Publics = () => {
  const [token, setToken] = useState("");
  const [userId, setUserId] = useState("");
  const [publications, setPublications] = useState([]);

  useEffect(() => {
    loadTokenFromCookie();
  }, []);

  const loadTokenFromCookie = () => {
    const userToken = Cookie.get('user');
    setToken(userToken);
  };

  useEffect(() => {
    fetch('http://localhost:8081/user', {
      headers: {
        'Authorization': 'Bearer ' + token
      }
    })
      .then(response => response.json())
      .then(data => {
        setUserId(data.id);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }, [token]);

  useEffect(() => {
    fetch(`http://localhost:8090/items/user/${userId}`)
      .then(response => response.json())
      .then(data => {
        setPublications(data);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }, [userId, token]);

  return (
    <div>
      <h1>User ID: {userId}</h1>
      <ul>
        {publications.map(publicationId => (
          <li key={publicationId}>{publicationId}</li>
        ))}
      </ul>
    </div>
  );
}

export default Publics;
