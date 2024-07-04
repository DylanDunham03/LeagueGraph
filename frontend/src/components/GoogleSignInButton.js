import React, { useEffect } from "react";

const GoogleSignInButton = ({ onSuccess, onError, children }) => {
  useEffect(() => {
    if (typeof window.google !== "undefined") {
      window.google.accounts.id.initialize({
        client_id: process.env.REACT_APP_GOOGLE_CLIENT_ID,
        callback: handleCredentialResponse,
      });
    }
  }, []);

  const handleCredentialResponse = (response) => {
    if (response.credential) {
      onSuccess(response.credential);
    } else {
      onError("No credential returned");
    }
  };

  const handleClick = () => {
    if (typeof window.google !== "undefined") {
      window.google.accounts.id.prompt();
    }
  };

  return <div onClick={handleClick}>{children}</div>;
};

export default GoogleSignInButton;
