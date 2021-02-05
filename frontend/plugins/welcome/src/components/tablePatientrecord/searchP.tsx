import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Table';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';
import { Avatar,TextField } from '@material-ui/core';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
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
import AddBoxIcon from '@material-ui/icons/AddBox';
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
 
  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ limit: 110, offset: 0 });
    setPatientrecord(res);
  };
  const handleChange = (event : any, value : unknown) => {
    setPat(value as string);
  };
  const sc = async () => {
    
setSe(Pat);
    var p = await http.listPatientrecord({ limit: 10, offset: 0 })
    let i = 0
    for (let u of p){
    if( u.name === Pat && u.id !== undefined)
    i = i+1
    }
    console.log("ผู้ป่วย = ", Pat)

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
    getPatientrecord();
    setLoading(false);
  }, [loading]);

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Patientrecord`}
>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาข้อมูลผู้ป่วย">
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  พบข้อมูลผู้ป่วย
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูลผู้ป่วย
                  </Alert>
                )}
            </div>
          ) : null}
            <Autocomplete
              id="patientname" 
              options={Patientrecord.map((option) => option.name)}
              onChange={handleChange}
              renderInput={(params) => (
                <TextField {...params} label="ชื่อผู้ป่วย" margin="normal" variant="outlined" style={{ width: "50ch"}} />
              )}
            />
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
            <Button variant="contained" color="primary" style={{backgroundColor: "#d500f9"}}>
              Home
           </Button>
          </Link>&emsp;
          <Link component={RouterLink} to="/createPatientrecord">
            <Button variant="contained" color="primary" style={{backgroundColor: "#9500ae"}} startIcon={<AddBoxIcon />} size="large">
              ลงทะเบียนผู้ป่วยนอก
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