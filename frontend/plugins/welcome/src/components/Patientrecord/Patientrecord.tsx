import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../tablePatientrecord';
import Button from '@material-ui/core/Button';
import { Typography, Grid, Avatar } from '@material-ui/core';
import AddBoxIcon from '@material-ui/icons/AddBox';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { useEffect } from 'react';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

import { Cookies } from 'react-cookie/cjs';//cookie

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

export default function Tables() {
const profile = { givenName: '' };
const http = new DefaultApi();

  const [loading, setLoading] = React.useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  useEffect(() => {
      const getImg = async () => {
          const res = await http.getUser({ id: Number(Img) });
          setLoading(false);
          setUsers(res);
      };
      getImg();
  }, [loading]);
 
 return (
   <Page theme={pageTheme.home}>
     <Header style={HeaderCustom} title={`Patientrecord`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
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