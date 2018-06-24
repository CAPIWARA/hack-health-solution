import HarshMenu from '@/components/Harsh/HarshMenu'
import HarshEntry from '@/components/Harsh/HarshEntry'
import HarshLabel from '@/components/Harsh/HarshLabel'
import HarshButton from '@/components/Harsh/HarshButton'
import HarshScreen from '@/components/Harsh/HarshScreen'
import HarshLayout from '@/components/Harsh/HarshLayout'
import HarshSelect from '@/components/Harsh/HarshSelect'
import HarshOptions from '@/components/Harsh/HarshOptions'

const components = {
  HarshMenu,
  HarshEntry,
  HarshLabel,
  HarshButton,
  HarshLayout,
  HarshScreen,
  HarshSelect,
  HarshOptions
};

function install (Vue) {
  Object.entries(components)
    .forEach(([ name, component ]) => Vue.component(name, component));
}

export default install;
