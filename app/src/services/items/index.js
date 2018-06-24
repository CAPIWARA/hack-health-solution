import Sarradinha from '@/assets/icons/Sarradinha.png';
import SarradinhaActive from '@/assets/icons/Sarradinha-Active.png';
import User from '@/assets/icons/User.png';
import UserActive from '@/assets/icons/User-Active.png';
import Calendar from '@/assets/icons/Calendar.png';
import CalendarActive from '@/assets/icons/Calendar-Active.png';
import Ellipsis from '@/assets/icons/Ellipsis.png';
import EllipsisActive from '@/assets/icons/Ellipsis-Active.png';

const items = [
  {
    name: 'Perfil',
    path: { name: 'Profile' },
    icon: User,
    iconActive: UserActive
  },
  {
    name: 'Sarradinha',
    path: { name: 'Sarradinha' },
    icon: Sarradinha,
    iconActive: SarradinhaActive
  },
  {
    name: 'Agênda',
    path: { name: 'Schedule' },
    icon: Calendar,
    iconActive: CalendarActive
  },
  {
    name: 'Mais',
    path: { name: 'Other' },
    icon: Ellipsis,
    iconActive: EllipsisActive
  }
];

export default items;
