import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from './Tables_ID';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useEffect } from 'react';
import { Avatar, TextField } from '@material-ui/core';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntPatientrights } from '../../api/models/EntPatientrights';
import Autocomplete from '@material-ui/lab/Autocomplete';
import Pagination from '@material-ui/lab/Pagination';
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
  const [Pat2, setPat2] = React.useState<string>();
  const [Sender, setSender] = React.useState<EntPatientrights[]>(Array);
  const classes = useStyles();
  const [alert1, setAlert1] = React.useState(true);
  const [alert2, setAlert2] = React.useState(true);
  const [loading, setLoading] = React.useState(true);
  const [status, setStatus] = React.useState(false);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);
  const [PatientrightsGet, setPatientrightsGet] = React.useState<EntPatientrights[]>([]);
  //------------------------
  const [page, setPage] = React.useState(1);
  const [numpage, setNumPage] = React.useState(1);
  var p = 5


  const getPage = async () => {
    const i = await http.patientrightsGet({ name: Pat2 });
    let j = i.length;
    let k = Math.ceil(j / p)
    console.log("i & k", i, " / ", k)
    setPatientrightsGet(i)
    setNumPage(k);
    
    validate(i);

    if (i.length != 0){
      setPage(1);
      let data = [i[0]]
      for (let _i_ = 1; _i_ < p; _i_++) {
        if (i[_i_] !== undefined) {
          data.push(i[_i_])
        }
      }
  
      console.log("data = ", data)
     setDataPage(data);
    }
  else{
    setDataPage(i);
  }



  };


  const handleChangePage = (event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
   
    if (PatientrightsGet.length != 0){
    let v = (value - 1) * p
    let data = [PatientrightsGet[0 + v]]
    for (let _i_ = 1; _i_ < p; _i_++) {
      if (PatientrightsGet[_i_ + v] !== undefined) {
        data.push(PatientrightsGet[_i_ + v])
      }
    }

    console.log("data = ", data)
   setDataPage(data);
  }
else{
  setDataPage( PatientrightsGet);
}

  };
  const setDataPage = (event: any) => {
    setSender(event);

  };
  //-------------------------


  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ limit: 110, offset: 0 });
    setPatientrecord(res);
  };
  const handleChange = (event: any, value: unknown) => {
    setPat(value as string);
  };


  const validate = (_j_:any) => {
    
    if (Pat2 === undefined || Pat2 === null || Pat2 === '') {

      setAlert2(false);//แสดงข้อมูลทั้งหมด
    } else {

      setAlert2(true);//ไม่พบสิทธ์
    }
    if (_j_.length > 0) {

      setStatus(true);
      setAlert1(true);//พบสิทธิ์

    } else {

      setStatus(true);
      setAlert1(false);//ใช้ Alert2

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
    getPage();
  }, [loading]);

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ยินดีต้อนรับ เข้าสู่ ระบบ ค้นหาเบียนสิทธิ์`}
        subtitle="ของโรงบาล">
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาสิทธิ์">

          {status ? (
            <div>
              {alert1 ? (
                alert2 ? (
                <Alert severity="success">
                  พบสิทธิ์
                </Alert>
                ) : (
                  <Alert severity="info" style={{ marginTop: 20 }}>
                    แสดงข้อมูลทั้งหมด
                  </Alert>
                )
              ) : (
                  alert2 ? (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่พบสิทธ์
                    </Alert>
                  ) : (
                      <Alert severity="error" style={{ marginTop: 20 }}>
                        ไม่มีข้อมูลในฐานข้อมูล
                      </Alert>
                    )
                )}
            </div>
          ) : null}

          <FormControl variant="outlined" className={classes.formControl}>
            <Autocomplete
              id="patientname"
              freeSolo
              options={Patientrecord.map((option) => option.name)}
              inputValue={Pat2}
              onChange={handleChange}
              onInputChange={(event, newInputValue) => {
                setPat2(newInputValue);
              }}
              renderInput={(params) => (
                <TextField {...params} label="ชื่อผู้ป่วย" margin="normal" variant="outlined" />
              )}
            />
          </FormControl>
          <Button
            onClick={() => {
              getPage();
            }
            }
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
        <div className={classes.root}>

          <ComponanceTable sim={Sender}></ComponanceTable>

        </div>
        <br />
        <div className={classes.root}>

          <Pagination count={numpage} page={page} boundaryCount={2} onChange={handleChangePage} />

        </div>

      </Content>
    </Page>
  );
};

export default Table;
