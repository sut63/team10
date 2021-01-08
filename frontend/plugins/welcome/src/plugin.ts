import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo'
import SignIn from './components/SignIn'

//createHistorytaking
import create_Patientrights from './components/create_patientrights';
import createHistorytaking from './components/createHistorytaking';
import CreateBill from './components/createBill'
import createTreatment from './components/treatment';
import CreateDoctorinfo from './components/Doctorinfo';
import createPatientrecord from './components/createPatientrecord';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/createBill', CreateBill);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/create_Patientrights', create_Patientrights);
    router.registerRoute('/createHistorytaking', createHistorytaking);
    router.registerRoute('/createTreatment', createTreatment);
    router.registerRoute('/createPatientrecord', createPatientrecord);
    router.registerRoute('/Doctorinfo', CreateDoctorinfo);
  },
});
