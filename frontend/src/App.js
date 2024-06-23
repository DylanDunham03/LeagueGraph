import React, { useEffect, useState, useRef } from 'react';
// src/App.js
import { ThemeProvider, useTheme } from './context/ThemeContext';
import { getPlayerGraph } from './api/PlayerGraphService';
import { renderNetwork } from './components/displayGraph';

const isDataFetched = false;

function AppContent() {
    const { theme, setTheme } = useTheme();
    const networkContainer = useRef(null);
    const [graphData, setGraphData] = useState(null);

    useEffect(() => {
      if (graphData == null) {  // Only fetch if graphData is not already loaded
        getPlayerGraph('americas')
          .then(data => {
            setGraphData(data);
            isDataFetched.current = true;  // Mark data as fetched
          })
          .catch(error => console.error('Error fetching graph data:', error));
      }
    }, [graphData]);

    useEffect(() => {
      if (networkContainer.current && graphData) {
          renderNetwork(networkContainer.current, graphData);
          // console.log(networkContainer.current);  // Log the current container (should be a div element)
          // console.log('Theme set to:', theme);  // Log the current class of the body
      }
  }, [graphData, theme]);

  useEffect(() => {
    document.body.className = theme;
    console.log('Body class updated to:', theme);  // Verify that the class is updated correctly
  }, [theme]);  // Ensure this runs every time the theme changes


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
