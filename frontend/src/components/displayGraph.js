import { DataSet, Network } from 'vis-network/standalone/esm/vis-network';

export function renderNetwork(container, graphData) {
    const nodes = new DataSet(
        graphData.playersList.map(player => ({
            id: player.puuid,
            label: player.riotidName,
            title: player.role,
            shape: 'dot',  // Gives a nice circular node
            size: 12,  // Size of the node
        }))
    );

    const edges = new DataSet(
        graphData.connectionsList.map(connection => ({
            from: connection.playeroneuuid,
            to: connection.playertwouuid,
            label: `Played ${connection.timesplayed} times`,
            length: 200,  // Controls the length of the spring in the graph
            width: 2,  // Thickness of the edge line
        }))
    );

    const data = {
        nodes,
        edges
    };

    const options = {
        interaction: { dragNodes: true },  // Enable node dragging
        physics: {
            enabled: true,
            barnesHut: {
                gravitationalConstant: -3000,
                centralGravity: 0.3,
                springLength: 95,
                springConstant: 0.04,
                damping: 0.09,
                avoidOverlap: 0.1
            },
            solver: 'barnesHut'
        },
        nodes: {
            borderWidth: 2,
            borderWidthSelected: 3,
        }
    };

    const network = new Network(container, data, options);
}
