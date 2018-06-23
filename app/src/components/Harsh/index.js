import HarshEntry from '@/components/Harsh/HarshEntry'
import HarshButton from '@/components/Harsh/HarshButton'
import HarshScreen from '@/components/Harsh/HarshScreen'

const components = {
  HarshEntry,
  HarshButton,
  HarshScreen
};

function install (Vue) {
  Object.entries(components)
    .forEach(([ name, component ]) => Vue.component(name, component));
}

export default install;
