import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'

import create_Patientrights from './components/create_patientrights';
import Historytaking from './components/Historytaking';
import tableHistorytaking from './components/tableHistorytaking';
import createHistorytaking from './components/createHistorytaking';
import CreateBill from './components/createBill'
import createTreatment from './components/createTreatment';
import Treatment from './components/Treatment';
import CreateDoctorinfo from './components/Doctorinfo';
import createPatientrecord from './components/createPatientrecord';
import Patientrecord from './components/Patientrecord';
import tablePatientrecord from './components/tablePatientrecord';
import Login from './components/Login';
import Logout from './components/Logout';
import reg from './components/reg';

import { Cookies } from 'react-cookie/cjs';//cookie
import React, { useEffect, FC } from 'react';




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
        router.registerRoute('/createBill', CreateBill);
      }
      if (Status == 'Med' || Status == 'Root' ) {
        router.registerRoute('/create_Patientrights', create_Patientrights);
        router.registerRoute('/createPatientrecord', createPatientrecord);
        router.registerRoute('/Patientrecord', Patientrecord);
        router.registerRoute('/tablePatientrecord', tablePatientrecord);
      }
      if (Status == 'Nur' || Status == 'Root' ) {
        router.registerRoute('/createHistorytaking', createHistorytaking);
        router.registerRoute('/Historytaking', Historytaking);
        router.registerRoute('/tableHistorytaking', tableHistorytaking);
      }
      if (Status == 'Doc' || Status == 'Root' ) {
        router.registerRoute('/Treatment', Treatment);
        router.registerRoute('/createTreatment', createTreatment);
        router.registerRoute('/tableHistorytaking', tableHistorytaking);

      }
      
      if (Status == 'Reg' || Status == 'Root' ) {
        router.registerRoute('/Doctorinfo', CreateDoctorinfo);
      }

    
      router.registerRoute('/', WelcomePage);
    }





  },
});
