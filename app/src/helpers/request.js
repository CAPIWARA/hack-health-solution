export const baseURL = 'http://localhost:3000/';

/**
 * Normaliza URLs removendo barras duplas no curpo da URL.
 */
export function normalizeURL (url) {
  const URL = /(https?:\/\/)?(.+)/;
  const [ , protocol = '', domainAndRoute = '' ] = URL.exec(url) || [];
  return protocol + domainAndRoute.replace(/\/+/g, '/');
}

/**
 * Obtém os cabeçalhos para a requisição.
 */
export function getHeaders () {
  const headers = {
    'Content-Type': 'application/json'
  };

  return headers;
}

/**
 * Realiza uma requisição e retorna a resposta.
 */
export async function request (method, route, params) {
  const url = normalizeURL(baseURL + '/' + route);
  const response = await fetch(url, {
    method,
    headers: getHeaders(params),
    body: JSON.stringify(params)
  });

  return response;
}

/**
 * Realiza uma requisição usando GraphQL e retorna os dados JSON de saída.
 */
export async function requestGraphQL (query, variables = {}) {
  const response = await request('POST', '/graphql', {
    query,
    variables
  });

  return response.json();
}
