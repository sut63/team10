import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';
import { Avatar,TextField } from '@material-ui/core';
import { EntDoctorinfo } from '../../api/models/EntDoctorinfo';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { Alert } from '@material-ui/lab';
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
  const http = new DefaultApi();
  const useStyles = makeStyles(theme => ({
    table: {
      minWidth: 650,
    },
    formControl: {
      margin: theme.spacing(3),
      width: 350,
    },
    root: {
      flexGrow: 1,
      display: 'flex',
      justifyContent: 'center',
    },
  }));
  const [Pat, setPat] = React.useState<string>();
  const [Se, setSe] = React.useState<string>();
  const classes = useStyles();
  const [alert, setAlert] = React.useState(true);
  const [loading, setLoading] = React.useState(true);
  const [status, setStatus] = React.useState(false);

  const [Users, setUsers] = React.useState<Partial<EntUser>>();
 
  const [Doctorinfo, setDoctorinfo] = React.useState<EntDoctorinfo[]>([]);

  const getDoctorinfo = async () => {
    const res = await http.listDoctorinfo({ limit: 110, offset: 0 });
    setDoctorinfo(res);
  };
  const handleChange = (event : any, value : unknown) => {
    setPat(value as string);
  };
  const sc = async () => {
    
setSe(Pat);
    var p = await http.listDoctorinfo({ limit: 1000000000, offset: 0 })
    let i = 0
    for (let u of p){
      
    if( u.licensenumber === Pat && u !== undefined)
    i = i+1
    
    }
    console.log("เเพทย์ = ", Pat)

    if (i != 0) {
      setStatus(true);
      setAlert(true);
    } else {
      setStatus(true);
      setAlert(false);
    }

    setTimeout(() => {
      setStatus(false);
    }, 5000);

    
  };




  useEffect(() => {
    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });

      setUsers(res);
    };
    getImg();
    getDoctorinfo();
    setLoading(false);
  }, [loading]);

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ยินดีต้อนรับเข้าสู่ระบบค้นหาข้อมูลเเพทย์`}
        subtitle="">
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาข้อมูลเเพทย์">

          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  พบข้อมูล
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูล
                  </Alert>
                )}
            </div>
          ) : null}

          <FormControl variant="outlined" className={classes.formControl}>
            <Autocomplete
              id="licensenumber"
             
              options={Doctorinfo.map((option) => option.licensenumber)}
              onChange={handleChange}
              renderInput={(params) => (
                <TextField {...params} label="เลขใบประกอบวิชาชีพ" margin="normal" variant="outlined" />
              )}
            />
            </FormControl>
            <Button
              onClick={() => {
               sc();
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

         <Link component={RouterLink} to="/Doctorinfo">
            <Button variant="contained" color="primary">
              เพิ่มข้อมูลเเพทย์
           </Button>
          </Link>


        </ContentHeader>
        <div className={classes.root}>

          <ComponanceTable sim={Se}></ComponanceTable>

        </div>

      </Content>
    </Page>
  );
};

export default Table;