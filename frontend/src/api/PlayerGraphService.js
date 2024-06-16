import {GraphRequest, PlayerGraphServiceClient} from '../protos/player_graph_service_grpc_web_pb';

const client = new PlayerGraphServiceClient('http://localhost:8080', null, null);

export const getPlayerGraph = async (region) => {
  const request = new GraphRequest();
  request.setRegion(region);

  return new Promise((resolve, reject) => {
    client.getPlayerGraph(request, {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response.toObject());
      }
    });
  });
};
