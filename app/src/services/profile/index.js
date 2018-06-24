import getProfile from '@/services/profile/getProfile.graphql';
import { requestGraphQL } from '@/helpers/request';

export async function fetch () {
  const response = await requestGraphQL(getProfile);
  const profile = response.data.user;
  return profile;
}
