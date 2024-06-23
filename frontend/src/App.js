import React, { useEffect, useState, useRef } from 'react';
import { getPlayerGraph } from './api/PlayerGraphService';
import { renderNetwork } from './components/displayGraph'; 

function App() {
  const [graphData, setGraphData] = useState(null);
  const networkContainer = useRef(null);
  const isDataFetched = useRef(false);

  useEffect(() => {
    if (!isDataFetched.current) {
      getPlayerGraph('americas')
        .then(data => {
          setGraphData(data);
          isDataFetched.current = true;  // Mark data as fetched
        })
        .catch(error => console.error('Error fetching graph data:', error));
    }
  }, []);

  useEffect(() => {
    if (networkContainer.current && graphData) {
      renderNetwork(networkContainer.current, graphData);
    }
  }, [graphData]);

  return (
    <div>
      <h1>Player Graph</h1>
      <div ref={networkContainer} style={{ height: '500px', width: '100%' }} />
    </div>
  );
}

export default App;
