import createSarrada from '@/services/sarrada/createSarrada.graphql';
import { requestGraphQL } from '@/helpers/request';

export async function create (params) {
  const response = await requestGraphQL(createSarrada, params);
  const sarrada = response.data.sarrada;
  return sarrada;
}
