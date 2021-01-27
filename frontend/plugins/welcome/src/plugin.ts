import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';


import create_Patientrights from './components/create_patientrights';
import Table_patientrights from './components/Table_patientrights';
import tableHistorytaking from './components/tableHistorytaking';
import createHistorytaking from './components/createHistorytaking';
import CreateBill from './components/createBill'
import BillSearch from './components/tableBill'
import createTreatment from './components/createTreatment';
import Treatment from './components/Treatment';
import findTreatment from './components/findTreatment';
import CreateDoctorinfo from './components/Doctorinfo';
import Table_Doctors from './components/Table_Doctors';
import create_Doctor from './components/create_Doctor';
import createPatientrecord from './components/createPatientrecord';
import tablePatientrecord from './components/tablePatientrecord';
import Login from './components/Login';
import Logout from './components/Logout';
import reg from './components/reg';

import { Cookies } from 'react-cookie/cjs';//cookie





const cookies = new Cookies();
const Status = cookies.get('Status');
const Log = cookies.get('Log');




export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if (Log == undefined) {
      router.registerRoute('/', Login);

    } else {
      router.registerRoute('/sign_out', Logout);

      if (Status == 'Root') {
        router.registerRoute('/reg', reg);
      }
      if (Status == 'Fin' || Status == 'Root' ) {
        router.registerRoute('/BillSearch', BillSearch);
        router.registerRoute('/createBill', CreateBill);
      }
      if (Status == 'Med' || Status == 'Root' ) {
        router.registerRoute('/create_Patientrights', create_Patientrights);
        router.registerRoute('/Table_patientrights', Table_patientrights);
        router.registerRoute('/createPatientrecord', createPatientrecord);
        router.registerRoute('/tablePatientrecord', tablePatientrecord);
      }
      if (Status == 'Nur' || Status == 'Root' ) {
        router.registerRoute('/createHistorytaking', createHistorytaking);
        router.registerRoute('/tableHistorytaking', tableHistorytaking);
      }
      if (Status == 'Doc' || Status == 'Root' ) {
        router.registerRoute('/Treatment', Treatment);
        router.registerRoute('/createTreatment', createTreatment);
        router.registerRoute('/findTreatment', findTreatment);
        router.registerRoute('/tableHistorytaking', tableHistorytaking);

      }
      
      if (Status == 'Reg' || Status == 'Root' ) {
        
        router.registerRoute('/Doctorinfo', CreateDoctorinfo);
        router.registerRoute('/Table_Doctors', Table_Doctors);
        router.registerRoute('/create_Doctor', create_Doctor);
      }

      router.registerRoute('/', WelcomePage);
    }





  },
});
