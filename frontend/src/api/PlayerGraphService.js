import {GraphRequest, PlayerGraphServiceClient} from '../protos/player_graph_service_grpc_web_pb';

const client = new PlayerGraphServiceClient('http://localhost:8080', null, null);

export const getPlayerGraph = async (region) => {
  const request = new GraphRequest();
  request.setRegion(region);

  console.log(`Requesting data for region: ${region}`);

  return new Promise((resolve, reject) => {
      client.getPlayerGraph(request, {}, (err, response) => {
          if (err) {
              console.error("Error in gRPC call:", err);
              reject(err);
          } else {
              console.log("Received response:", response.toObject());
              resolve(response.toObject());
          }
      });
  });
};

