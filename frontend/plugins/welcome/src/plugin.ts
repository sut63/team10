import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'

//createHistorytaking
import create_Patientrights from './components/create_patientrights';
import createHistorytaking from './components/createHistorytaking';
import CreateBill from './components/createBill'
import createTreatment from './components/treatment';
import CreateDoctorinfo from './components/Doctorinfo';
import createPatientrecord from './components/createPatientrecord';
import Login from './components/Login';
import reg from './components/reg';

import { Cookies } from 'react-cookie/cjs';//cookie
import React, { useEffect, FC } from 'react';




const cookies = new Cookies();
const Status = cookies.get('Status');




export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/', Login);
    //if (Status == 'Root') {
    router.registerRoute('/reg', reg);
    //}
    if (Status == 'Fin') {
      router.registerRoute('/createBill', CreateBill);
    }
    if (Status == 'Med') {
      router.registerRoute('/create_Patientrights', create_Patientrights);
    }
    if (Status =='Nur') {
      router.registerRoute('/createHistorytaking', createHistorytaking);
    }
    if (Status =='Doc') {
      router.registerRoute('/createTreatment', createTreatment);
    }
    if (Status =='Med') {
      router.registerRoute('/createPatientrecord', createPatientrecord);
    }
    if (Status =='Reg') {
      router.registerRoute('/Doctorinfo', CreateDoctorinfo);
    }
  },
});
