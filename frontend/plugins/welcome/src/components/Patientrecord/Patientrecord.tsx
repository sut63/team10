import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../tablePatientrecord';
import Button from '@material-ui/core/Button';
import { Typography, Grid } from '@material-ui/core';
import AddBoxIcon from '@material-ui/icons/AddBox';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
export default function Tables() {
const profile = { givenName: '' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
      title={`${profile.givenName || 'Patientrecord'}`}
     ></Header>
     <Content>
       <ContentHeader title="ข้อมูลผู้ป่วยนอก">
       </ContentHeader>
       <ContentHeader title="">
       <Grid item  xs={12} md={12}>
       <Typography variant="h6" gutterBottom align="center">
         <div>
          <Link component={RouterLink} to="/createPatientrecord">
            <Button variant="contained" color="primary" style={{backgroundColor: "#b388ff"}} startIcon={<AddBoxIcon />} size="large">
              ลงทะเบียนผู้ป่วยนอก
            </Button>
          </Link>
         </div>
      </Typography>
      </Grid>
      </ContentHeader>
      <div>
      </div>
      <ComponanceTable></ComponanceTable>
    </Content>
  </Page>
 );
};