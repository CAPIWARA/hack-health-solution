import createSarrada from '@/services/sarrada/createSarrada.graphql';
import getSarrada from '@/services/sarrada/getSarrada.graphql';
import getSarradas from '@/services/sarrada/getSarradas.graphql';
import { requestGraphQL } from '@/helpers/request';

export async function fetch (id) {
  const response = await requestGraphQL(getSarrada, { id });
  const sarrada = response.data.sarrada;
  return sarrada;
}

export async function create (params) {
  const response = await requestGraphQL(createSarrada, params);
  const sarrada = response.data.sarrada;
  return sarrada;
}

export async function getHistory () {
  const response = await requestGraphQL(getSarradas);
  return response.data;
}
