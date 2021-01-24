import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

import {

  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Avatar,

  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { ControllersPatientrights } from '../../api/models/ControllersPatientrights';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntAbilitypatientrights } from '../../api/models/EntAbilitypatientrights';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntMedicalrecordstaff } from '../../api/models/EntMedicalrecordstaff';
import { Alert } from '@material-ui/lab';
import { Cookies } from 'react-cookie/cjs';//cookie
import { EntUser } from '../../api/models/EntUser';


import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2';


// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const Name = cookies.get('Name');

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: '25ch',
  },


  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },

}));


const NewPatientright: FC<{}> = () => {
  const classes = useStyles();
 
  const http = new DefaultApi();
  const cookies = new Cookies();

  const [Patientrights, setPatientrights] = React.useState<Partial<ControllersPatientrights>>({});

  const [Abilitypatientrights, setAbilitypatientrights] = React.useState<EntAbilitypatientrights[]>([]);
  const [Patientrecord, setPatientrecord] = React.useState<EntPatientrecord[]>([]);

  const [Insurance, setInsurance] = React.useState<EntInsurance[]>([]);
  const [Medicalrecordstaff, setMedicalrecordstaff] = React.useState<Partial<EntMedicalrecordstaff>>();
  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  const med = cookies.get('Med');
  const Img = cookies.get('Img');


  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 5000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getChangeOfUser = async () => {

    const name = "medicalrecordstaff";
    const value = parseInt(med, 10);
    setPatientrights({ ...Patientrights, [name]: value });
  };



  const getMedicalrecordstaffs = async () => {

    const res = await http.getMedicalrecordstaff({ id: Number(med) });
    setMedicalrecordstaff(res);
  };

  const getAbilitypatientrights = async () => {
    const res = await http.listAbilitypatientrights({ limit: 10, offset: 0 });
    setAbilitypatientrights(res);
  };

  const getPatientrecord = async () => {
    const res = await http.listPatientrecord({ limit: 10, offset: 0 });
    setPatientrecord(res);
  };

  const getInsurance = async () => {
    const res = await http.listInsurance({ limit: 10, offset: 0 });
    setInsurance(res);
  };

  const getImg = async () => {
    const res = await http.getUser({ id: Number(Img) });
    setUsers(res);
  };



  // Lifecycle Hooks
  useEffect(() => {
    getChangeOfUser();
    getMedicalrecordstaffs();
    getAbilitypatientrights();
    getPatientrecord();
    getInsurance();
    getChangeOfUser();
    getImg();
  }, []);

  // set data to object Patientright
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof NewPatientright;
    const { value } = event.target;

    setPatientrights({ ...Patientrights, [name]: value });
  };


  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }


  const CreatePatientright = async () => {

    const apiUrl = 'http://localhost:8080/api/v1/patientrightss';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Patientrights),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          alertMessage("error", data.error);
        }
      });
  };


  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ลงทะเบียนสิทธิ์`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <ContentHeader title="ข้อมูล">

          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    บันทึกไม่สำเร็จ
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <TextField
                name="permission"
                label="ชื่อสิทธิ์"
                variant="outlined"
                type="string"
                size="medium"

                value={Patientrights.permission}
                onChange={handleChange}
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <TextField
                name="permissionArea"
                label="พื้นที่ที่ใช้สิทธิ์ได้"
                variant="outlined"
                type="string"
                size="medium"

                value={Patientrights.permissionArea}
                onChange={handleChange}
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <TextField
                name="responsible"
                label="ผู้รับผิดชอบดูแลสิทธิ์"
                variant="outlined"
                type="string"
                size="medium"

                value={Patientrights.responsible}
                onChange={handleChange}
              />
            </FormControl>
          </form>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>ผู้ป่วย</InputLabel>
              <Select
                name="patientrecord"
                value={Patientrights.patientrecord}
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
          </form>
        </div>


        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>ประกัน</InputLabel>
              <Select
                name="insurance"
                value={Patientrights.insurance}
                onChange={handleChange}
              >
                {Insurance.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.insurancecompany}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </form>
        </div>


        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>รูปแบบสิทธิ์</InputLabel>
              <Select
                name="abilitypatientrights"
                value={Patientrights.abilitypatientrights}
                onChange={handleChange}
              >
                {Abilitypatientrights.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      ตรวจสุขภาพ และ ค่า แลป : {item.examine} หัตถการ  : {item.operative} เวชภัณฑ์  : {item.medicalSupplies}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </form>
        </div>



        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <TextField
                disabled
                id="outlined-disabled"
                label="พนักงานเวชระเบียง"
                defaultValue=" "
                value={Medicalrecordstaff?.name as string}
                variant="outlined"
              />
            </FormControl>
          </form>
        </div>


        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreatePatientright();
                }}
                variant="contained"
                color="primary"
              >
                Submit
             </Button>
              <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                Back
                </Button>
              <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/Table_patientrights"
                variant="contained"
                color="secondary"
              >
                SHOW
                </Button>
            </div>
          </form>


        </div>
      </Content>
    </Page>
  );
};
export default NewPatientright;