import React, { useEffect, useState } from 'react';
import { getPlayerGraph } from './api/PlayerGraphService';

function App() {
  const [graphData, setGraphData] = useState(null);

  useEffect(() => {
    getPlayerGraph('americas')
      .then(data => {
        setGraphData(data);
        console.log(data);
      })
      .catch(error => {
        console.error('Error fetching graph data:', error);
      });
  }, []);

  return (
    <div>
      <h1>Player Graph</h1>
      {/* Render your graph data here */}
    </div>
  );
}

export default App;
