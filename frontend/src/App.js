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
        <div className={`${theme} flex flex-col items-center justify-center min-h-screen`}>
            <div className="absolute top-0 right-0 p-4">
                <button onClick={toggleTheme} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Toggle Theme
                </button>
            </div>
            <h1 className="text-4xl mb-4">Player Graph</h1>
            <div ref={networkContainer} style={{ height: '500px', width: '100%' }} />
            <button className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded">
                Sign In
            </button>
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
