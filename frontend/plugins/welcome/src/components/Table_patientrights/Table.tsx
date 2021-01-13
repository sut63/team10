import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';
import { Avatar } from '@material-ui/core';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const Table: FC<{}> = () => {
 const profile = { givenName: 'ระบบ ลงทะเบียนสิทธิ์' };
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
     <Header
       title={`ยินดีต้อนรับ เข้าสู่ ${profile.givenName || 'to Backstage'}`}
       subtitle="ของโรงบาล">  
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>    
       </Header>
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