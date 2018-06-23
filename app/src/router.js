import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/screens/Home')
    },
    {
      path: '/entrar',
      name: 'Login',
      component: () => import('@/screens/Login')
    },
    {
      path: '/inscrever-se',
      name: 'Subscribe',
      component: () => import('@/screens/Subscribe')
    }
  ]
});

export default router;
