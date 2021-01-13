import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables';
import Button from '@material-ui/core/Button';
import Timer from '../Timer';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const Table: FC<{}> = () => {
 const profile = { givenName: 'ระบบ ลงทะเบียนสิทธิ์' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ยินดีต้อนรับ เข้าสู่ ${profile.givenName || 'to Backstage'}`}
       subtitle="ของโรงบาล"
      >      <Timer /></Header>
           <Content>
       <ContentHeader title="ลงทะเบียนสิทธิ์">
         
          
         
         
         <Link component={RouterLink} to="/">
           <Button variant="contained" color="primary">
           Home
           </Button>
         </Link>&emsp;
         
         <Link component={RouterLink} to="/create_Patientrights">
           <Button variant="contained" color="primary">
           ลงทะเบียนสิทธิ์
           </Button>
         </Link>


       </ContentHeader>
       
       <ComponanceTable></ComponanceTable>
        
       
     </Content>
   </Page>
 );
};
 
export default Table;