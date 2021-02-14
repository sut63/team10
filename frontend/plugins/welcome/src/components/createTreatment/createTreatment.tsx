import React, { useEffect, FC } from 'react';
import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, MenuItem, Button, TextField, FormControl, Select, Link, Typography } from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';

import { Alert, AlertTitle } from '@material-ui/lab';
import Swal from 'sweetalert2'; // alert
import { Avatar } from '@material-ui/core';
import { Cookies } from 'react-cookie/cjs';//cookie


import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntTypetreatment } from '../../api/models/EntTypetreatment';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntTreatment } from '../../api/models/EntTreatment';

import { ControllersTreatment } from '../../api/models/ControllersTreatment';

import Treatment from '../Treatment';

const cookies = new Cookies();
const Name = cookies.get('Name');
const Doc = cookies.get('Doc');
const Img = cookies.get('Img');

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    button: {
      display: 'block',
      marginTop: theme.spacing(2),
    },
    formControl: {
      width: 400,
    },
    selectEmpty: {
      marginTop: theme.spacing(1),
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    table: {
      minWidth: 650,
    },
  }),
);

// header css
const HeaderCustom = {
  minHeight: '50px',
};
interface Treatments {
  symptom: string,
  treat: string,
  medicine: string,
  typetreatment: number,
  doctor: number,
  patientrecord: number,
}

const createTreatment: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = React.useState(false);
  const [loading, setLoading] = React.useState(true);

  const [msg, setAlert] = React.useState(String);
  const [title, setAlerttitle] = React.useState(String);

  const [symptomError, setsymptomError] = React.useState('');
  const [treatError, settreatError] = React.useState('');
  const [medicineError, setmedicineError] = React.useState('');

  const [treatment, setTreatment] = React.useState<Partial<ControllersTreatment>>({});

  const [typetreatments, setTypetreatments] = React.useState<EntTypetreatment[]>([]);
  const [doctors, setDoctors] = React.useState<Partial<EntDoctor>>();
  const [patientrecords, setPatientrecords] = React.useState<EntPatientrecord[]>([]);
  const [treatments, setTreatments] = React.useState<EntTreatment[]>([]);

  const [treates, settreat] = React.useState(String);
  const [symptomes, setsymptom] = React.useState(String);
  const [medicinees, setmedicine] = React.useState(String);
  const [datetreat, setDatetreat] = React.useState(String);
  const [doctorid, setdoctorId] = React.useState(Number);
  const [patientrecordid, setpatientrecordId] = React.useState(Number);
  const [typetreatmentid, settypetreatmentId] = React.useState(Number);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

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

  useEffect(() => {
    const getDocdor = async () => {
      const res = await http.getDoctor({ id: Number(Doc) });
      setLoading(false);
      setDoctors(res);
    };
    const getTypetreatment = async () => {
      const res = await http.listTypetreatment({ offset: 0 });
      setLoading(false);
      setTypetreatments(res);
    };
    const getPatientrecord = async () => {
      const res = await http.listPatientrecord({ offset: 0 });
      setLoading(false);
      setPatientrecords(res);
    };
    const getTreatment = async () => {
      const res = await http.listTreatment({ offset: 0 });
      setLoading(false);
      setTreatments(res);
    };
    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getPatientrecord();
    getTypetreatment();
    getTreatment();
    getDocdor();
    getImg();
  }, [loading]);

  const refreshPage = () => {
    window.location.reload();
  }
  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof Create_Treatment;
    const { value } = event.target;
    const validateValue = value.toString()
    //checkPattern(name, validateValue)
    setTreatment({ ...treatment, [name]: value, doctor: doctors?.id });
  };

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch (field) {
      case 'Patientrecord not found':
        alertMessage("error", "กรุณาเลือกผู้ป่วย");
        return;
      case 'Typetreatment not found':
        alertMessage("error", "กรุณาเลือกรูปแบบการรักษา");
        return;
      case 'Symptom ไม่ถูกต้อง กรุณากรอกเป็นภาษาไทย':
        alertMessage("error", "กรุณากรอกอาการต้องเป็นภาษาไทย");
        return;
      case 'Treat ไม่ถูกต้อง กรุณากรอกเป็นภาษาไทย':
        alertMessage("error", "กรุณากรอกรายละเอียดการรักษาต้องเป็นภาษาไทย");
        return;
      case 'Medicine ไม่ถูกต้อง กรุณากรอกเป็นภาษาไทย':
        alertMessage("error", "กรุณากรอกชื่อยาต้องเป็นภาษาไทย");
        return;
      case 'ent: validator failed for field "Symptom": value does not match validation':
        alertMessage("error", "อาการต้องเป็นภาษาไทย");
        return;
      case 'ent: validator failed for field "Treat": value does not match validation':
        alertMessage("error", "รายละเอียดการรักษาต้องเป็นภาษาไทย");
        return;
      case 'ent: validator failed for field "Medicine": value does not match validation':
        alertMessage("error", "ชื่อยาต้องเป็นภาษาไทย");
        return;
      default:
        alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  const Create_Treatment = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/Treatments';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(treatment),
    };
    console.log(treatment);
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(async data => {
        console.log(data.error);
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error);
        }
      });
  };

  return (
    <div>
      <Page theme={pageTheme.home}>
        <Header style={HeaderCustom} title={`Treatment Department`}>
          <Avatar alt="Remy Sharp" src={Users?.images as string} />
          <div style={{ marginLeft: 10 }}>{Name}</div>
        </Header>
        <Content>
          <Grid container item xs={12} sm={12}  >
            <Grid item xs={12}>
              <Typography align="center">
                <Typography align="center" variant="h3">
                  <br />----  Create Treatment  ----
                    </Typography>
                <form noValidate autoComplete="off">
                  <FormControl className={classes.formControl}>
                    <Typography align="center" variant="h6">
                      <br />แพทย์
                    <br />{doctors?.edges?.edgesOfDoctorinfo?.doctorname} {doctors?.edges?.edgesOfDoctorinfo?.doctorsurname}
                    </Typography>

                  </FormControl>
                </form>
                <br />
                <form noValidate autoComplete="off">
                  <FormControl className={classes.formControl}>
                    <Typography align="center" variant="h6">
                      <br />ผู้ป่วย
                        </Typography>
                    <Select
                      name="patientrecord"
                      value={treatment.patientrecord}
                      onChange={handleChange}
                    >
                      {patientrecords.map(item => {
                        return (
                          <MenuItem value={item.id}>
                            {item.name}
                          </MenuItem>
                        );
                      })}
                    </Select>
                  </FormControl>
                </form>
                <br />
                <form noValidate autoComplete="off">
                  <FormControl className={classes.formControl}>
                    <Typography align="center" variant="h6">
                      <br />รูปแบบการรักษา
                        </Typography>
                    <Select
                      name="Typetreatment"
                      value={treatment.typetreatment}
                      onChange={handleChange}
                    >
                      {typetreatments.map(item => {
                        return (
                          <MenuItem value={item.id}>
                            {item.typetreatment}
                          </MenuItem>
                        );
                      })}
                    </Select>
                  </FormControl>
                </form>
                <form noValidate autoComplete="off">
                  <FormControl className={classes.formControl}>
                    <Typography align="center" variant="h6">
                      <br />อาการ<br />
                      <TextField
                        name="Symptom"
                        id="Symptom"
                        className={classes.formControl}
                        value={treatment.symptom}
                        multiline
                        onChange={handleChange} />
                    </Typography>
                  </FormControl>
                </form>
                <br />
                <form noValidate autoComplete="off">
                  <FormControl className={classes.formControl}>
                    <Typography align="center" variant="h6">
                      <br />รายละเอียดการรักษา<br />
                      <TextField
                        name="Treat"
                        id="Treat"
                        className={classes.formControl}
                        value={treatment.treat}
                        multiline
                        onChange={handleChange} />
                    </Typography>
                  </FormControl>
                </form>
                <br />
                <form noValidate autoComplete="off">
                  <FormControl className={classes.formControl}>
                    <Typography align="center" variant="h6">
                      <br />ยารักษา<br />
                      <TextField
                        name="Medicine"
                        id="Medicine"
                        className={classes.formControl}
                        value={treatment.medicine}
                        multiline
                        onChange={handleChange} />
                    </Typography>
                  </FormControl>
                </form>
                <br />

                <Typography align="center">
                  <br />
                  <Button
                    onClick={() => {
                      Create_Treatment();
                    }}

                    variant="contained"
                    color="primary"
                  >
                    บันทึกการรักษา
                </Button>
                </Typography>
                <br />
                <Typography align="center">
                  <Link component={RouterLink} to="/findTreatment">
                    <Button variant="contained" color="primary" style={{ backgroundColor: "#21b6ae" }}>
                      ย้อนกลับ
            </Button>
                  </Link>
                </Typography>
              </Typography>
              <br />
            </Grid>
          </Grid>
        </Content>
      </Page>
    </div>
  );
};

export default createTreatment;