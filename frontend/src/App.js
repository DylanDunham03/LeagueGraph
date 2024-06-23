import React, { useEffect, useState, useRef } from 'react';
// src/App.js
import { ThemeProvider, useTheme } from './context/ThemeContext';
import { getPlayerGraph } from './api/PlayerGraphService';
import { renderNetwork } from './components/displayGraph';
import { calculateLocalStorageSize } from './components/utils';

function AppContent() {
    const { theme, setTheme } = useTheme();
    const networkContainer = useRef(null);
    const [graphData, setGraphData] = useState(null);

    // useEffect(() => {
    //   if (graphData == null) {  // Only fetch if graphData is not already loaded
    //     getPlayerGraph('americas')
    //       .then(data => {
    //         setGraphData(data);
    //         isDataFetched.current = true;  // Mark data as fetched
    //       })
    //       .catch(error => console.error('Error fetching graph data:', error));
    //   }
    // }, [graphData]);
    useEffect(() => {
      // Attempt to load graph data from localStorage first
      const savedGraphData = localStorage.getItem('graphData');
      if (savedGraphData) {
          setGraphData(JSON.parse(savedGraphData)); // Parse the string back to an object
      } else {
          // Fetch data only if it's not found in local storage
          getPlayerGraph('americas')
              .then(data => {
                  setGraphData(data);
                  localStorage.setItem('graphData', JSON.stringify(data)); // Save to localStorage
              })
              .catch(error => console.error('Error fetching graph data:', error));
      }
  }, []); // Empty dependency array ensures this effect runs only once after the component mounts


    useEffect(() => {
      if (networkContainer.current && graphData) {
          renderNetwork(networkContainer.current, graphData);
      }
  }, [graphData]);

  useEffect(() => {
    document.body.className = theme;
    console.log('Body class updated to:', theme);  // Verify that the class is updated correctly
  }, [theme]);  // Ensure this runs every time the theme changes

  useEffect(() => {
    const localStorageSize = calculateLocalStorageSize();
    console.log('Local Storage size in KB:', localStorageSize);
}, []);


    const toggleTheme = () => {
        setTheme(theme === 'light' ? 'dark' : 'light');
    };

    return (
        <div className={theme}>
            <h1 style={{ textAlign: 'center' }}>Player Graph</h1>
            <button onClick={toggleTheme}>Toggle Theme</button>
            <div ref={networkContainer} style={{ height: '500px', width: '100%' }} />
        </div>
    );
}

function App() {
    return (
        <ThemeProvider>
            <AppContent />
        </ThemeProvider>
    );
}

export default App;
