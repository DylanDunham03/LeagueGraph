import React, { useEffect, useState } from 'react';
import { getPlayerGraph } from './api/PlayerGraphService';

function App() {
  const [graphData, setGraphData] = useState(null);

  console.log(`Starting getPlayerGraph`);
  useEffect(() => {
    getPlayerGraph('americas')
      .then(data => {
        setGraphData(data);
        console.log(data);
      })
      .catch(error => {
        console.error('Error fetching graph data:', error);
      });
      console.log(`Ending getPlayerGraph`);
  }, []);

  return (
    <div>
      <h1>Player Graph</h1>
      {/* Render your graph data here */}
    </div>
  );
}

export default App;
