import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';
import { Avatar } from '@material-ui/core';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

const cookies = new Cookies();
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const Table: FC<{}> = () => {
  const profile = { givenName: 'ระบบ ลงทะเบียนสิทธิ์' };
  const http = new DefaultApi();
  const useStyles = makeStyles(theme => ({
    table: {
      minWidth: 650,
    },
    formControl: {
      margin: theme.spacing(3),
      width: 350,
    },
  }));
  const [Pat, setPat] = React.useState<number>(0);
  const classes = useStyles();
  const [loading, setLoading] = React.useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ limit: 110, offset: 0 });
    setPatientrecord(res);
  };
  const handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPat(event.target.value as number);

  };

  useEffect(() => {
    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getImg();
    getPatientrecord();
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
          <FormControl variant="outlined" className={classes.formControl}>
            <InputLabel>ผู้ป่วย</InputLabel>
            <Select
              name="patientrecord"
              value={Pat}
              onChange={handleChange}
            >
              {Patientrecord.map((item: any) => {
                return (
                  <MenuItem key={item.id} value={item.id}>
                    {item.name}
                  </MenuItem>
                );
              })}
            </Select>
          </FormControl>
          <Button
            onClick={() => {
              //Sc();
            }}
            style={{ marginLeft: 10 }}
            variant="contained"
            color="primary"
          >
            ค้นหา
               </Button>&emsp;
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

        <ComponanceTable sim={Pat}></ComponanceTable>


      </Content>
    </Page>
  );
};

export default Table;