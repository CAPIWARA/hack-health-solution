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
    },
    {
      path: '/sarrada',
      component: () => import('@/screens/Sarrada'),
      children: [
        {
          path: '',
          name: 'Criar Sarrada',
          component: () => import('@/screens/Sarrada/Create')
        }
      ]
    },
    {
      path: '/perfil',
      name: 'Profile',
      component: () => import('@/screens/Profile')
    }
  ]
});

export default router;
