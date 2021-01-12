import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../tableHistorytaking';
import Button from '@material-ui/core/Button';
import { Typography, Grid, Avatar } from '@material-ui/core';
 
import { Cookies } from 'react-cookie/cjs';//cookie

import { Image5Base64Function } from '../../image/Image5';

import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');

export default function HistorytakingTable() {
const profile = { givenName: '' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header style={HeaderCustom} title={`HISTORYTAKING DEPARTMENT`}>
        <Avatar alt="Remy Sharp" src={Image5Base64Function} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
     <Content>
       <ContentHeader title="">
       <Grid item  xs={12} md={12}>
       <Typography variant="h6" gutterBottom align="center">
         <div>
          <Link component={RouterLink} to="/createHistorytaking">
            <Button variant="contained" color="primary" style={{backgroundColor: "#21b6ae"}}>
              ADD DATA
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

