import React, { useEffect, useState, useRef } from "react";
// src/App.js
import { ThemeProvider, useTheme } from "./context/ThemeContext";
import { getPlayerGraph } from "./api/PlayerGraphService";
import { renderNetwork } from "./components/displayGraph";
import { calculateLocalStorageSize } from "./components/utils";

//Styling
import "./styles.css";

const style = {
  wrapper: "flex min-h-screen flex-col items-center justify-center space-y-2",
  loginBtn:
    "group flex items-center space-x-4 rounded border-gray-300 border p-4 hover:bg-white",
  loginBtnText: "font-bold group-hover:text-black",
  text: "font-medium text-7xl",
  spacer: "py-8",
};

function Login() {
  const { theme, setTheme } = useTheme();
  const networkContainer = useRef(null);
  const [graphData, setGraphData] = useState(null);

  useEffect(() => {
    // Attempt to load graph data from localStorage first
    const savedGraphData = localStorage.getItem("graphData");
    if (savedGraphData) {
      setGraphData(JSON.parse(savedGraphData)); // Parse the string back to an object
    } else {
      // Fetch data only if it's not found in local storage
      getPlayerGraph("americas")
        .then((data) => {
          setGraphData(data);
          localStorage.setItem("graphData", JSON.stringify(data)); // Save to localStorage
        })
        .catch((error) => console.error("Error fetching graph data:", error));
    }
  }, []); // Empty dependency array ensures this effect runs only once after the component mounts

  useEffect(() => {
    if (networkContainer.current && graphData) {
      renderNetwork(networkContainer.current, graphData);
    }
  }, [graphData]);

  useEffect(() => {
    document.body.className = theme;
    console.log("Body class updated to:", theme); // Verify that the class is updated correctly
  }, [theme]); // Ensure this runs every time the theme changes

  useEffect(() => {
    const localStorageSize = calculateLocalStorageSize();
    console.log("Local Storage size in KB:", localStorageSize);
  }, []);

  const toggleTheme = () => {
    setTheme(theme === "light" ? "dark" : "light");
  };
  const signInWithGoogle = () => {
    console.log("Sign in with Google triggered"); // Replace with your auth logic
  };

  return (
    <div
      className={`${theme} flex flex-col items-center justify-center min-h-screen`}
    >
      <div className="absolute top-0 right-0 p-4 flex space-x-4">
        <button onClick={toggleTheme} className="btn-green">
          Toggle Theme
        </button>
        <button onClick={signInWithGoogle} className="btn-blue">
          Sign In
        </button>
      </div>
      <h1 className={`txt-general ${theme === 'dark' ? 'text-white' : 'text-black'}`}>
        League Friend Graph
      </h1>
      <div ref={networkContainer} style={{ height: "500px", width: "100%" }} />
      <button
        onClick={signInWithGoogle}
        //   className={style.loginBtn}
        className="btn-blue"
      >
        <img
          className="h-6 w-6"
          src="https://cdn-icons-png.flaticon.com/512/2991/2991148.png"
        />
        <span className={style.loginBtnText}>Sign in with Google</span>
      </button>
    </div>
  );
}

function App() {
  return (
    <ThemeProvider>
      <Login />
    </ThemeProvider>
  );
}

export default App;
