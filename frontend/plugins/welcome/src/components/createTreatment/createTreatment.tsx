import React, { useEffect, FC } from 'react';
import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, MenuItem, Button, TextField, FormControl, Select, Link, Typography } from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';

import { Alert } from '@material-ui/lab';
import Paper from '@material-ui/core/Paper';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Avatar } from '@material-ui/core';
import { Cookies } from 'react-cookie/cjs';//cookie


import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntTypetreatment } from '../../api/models/EntTypetreatment';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntTreatment } from '../../api/models/EntTreatment';





// header css
const HeaderCustom = {
  minHeight: '50px',
};

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
const createTreatment: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);
  const [loading, setLoading] = React.useState(true);

  const [typetreatments, setTypetreatments] = React.useState<EntTypetreatment[]>([]);
  const [doctors, setDoctors] = React.useState<Partial<EntDoctor>>();
  const [patientrecords, setPatientrecords] = React.useState<EntPatientrecord[]>([]);
  const [treatments, setTreatments] = React.useState<EntTreatment[]>([]);

  const [treatmentes, settreatment] = React.useState(String);
  const [datetreat, setDatetreat] = React.useState(String);
  const [doctorid, setdoctorId] = React.useState(Number);
  const [patientrecordid, setpatientrecordId] = React.useState(Number);
  const [typetreatmentid, settypetreatmentId] = React.useState(Number);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  useEffect(() => {
    const getDocdor = async () => {
      const res = await http.getDoctor({ id: Number(Doc) });
      setLoading(false);
      setDoctors(res);
    };
    const getTypetreatment = async () => {
      const res = await http.listTypetreatment({ limit: 3, offset: 0 });
      setLoading(false);
      setTypetreatments(res);
    };
    const getPatientrecord = async () => {
      const res = await http.listPatientrecord({ limit: 2, offset: 0 });
      setLoading(false);
      setPatientrecords(res);
    };
    const getTreatment = async () => {
      const res = await http.listTreatment({ limit: 100, offset: 0 });
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

  const TypetreatmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    settypetreatmentId(event.target.value as number);
  };
  const DoctorhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setdoctorId(event.target.value as number);
  };
  const TreatmenthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    settreatment(event.target.value as string);
  };
  const PatientrecordhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpatientrecordId(event.target.value as number);
  };
  const handleDatetimeChange = (event: any) => {
    setDatetreat(event.target.value as string);
  };

  const createTreatment = async () => {
  if ((treatmentes != '') 
      && (typetreatmentid != null) && (patientrecordid != null)) {
  
    const tm = {
      treatment: treatmentes,
      datetreat: datetreat,
      typetreatment: typetreatmentid,
      doctor: Number(Doc),
      patientrecord: patientrecordid,
    };
    console.log(tm);
    const res: any = await http.createTreatment({ treatment: tm });
        
    if (res.id != '') {
      setStatus(true);
      setAlert(true);
      setTimeout(() => {
        setStatus(false);
      }, 5000);
    }

  }
  else {
    setStatus(true);
    setAlert(false);
    setTimeout(() => {
      setStatus(false);
    }, 5000);
  }
};

  return (
    <div>
      {status ? (
        <div>
          {alert ? (
            <Alert severity="success">
              บันทึกการรักษาสำเร็จ
            </Alert>
          ) : (
              <Alert severity="warning" style={{ marginTop: 20 }}>
                บันทึกการรักษาไม่สำเร็จ
              </Alert>
            )}
        </div>
      ) : null}


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
                  <br />----  Create Traetment  ----
                    </Typography>

                <FormControl className={classes.formControl}>
                  <Typography align="center" variant="h6">
                    <br />แพทย์
                    <br />{doctors?.edges?.doctorinfo?.doctorname} {doctors?.edges?.doctorinfo?.doctorsurname}
                  </Typography>

                </FormControl>
                <br />
                <FormControl className={classes.formControl}>
                  <Typography align="center" variant="h6">
                    <br />ผู้ป่วย
                        </Typography>
                  <Select
                    name="patientrecord"
                    value={patientrecordid}
                    onChange={PatientrecordhandleChange}
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
                <br />
                <FormControl className={classes.formControl}>
                  <Typography align="center" variant="h6">
                    <br />รูปแบบการรักษา
                        </Typography>
                  <Select
                    name="typetreatment"
                    value={typetreatmentid}
                    onChange={TypetreatmenthandleChange}
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
                <Typography align="center" variant="h6">
                  <br />รายละเอียดการรักษา<br />
                  <TextField
                    className={classes.formControl}
                    value={treatmentes}
                    multiline
                    onChange={TreatmenthandleChange} />
                </Typography>



                <br />

                <Typography align="center">
                  <br />
                  <Button
                    onClick={() => {
                      createTreatment();
                    }}

                    variant="contained"
                    color="primary"
                  >
                    บันทึกการรักษา
                </Button>
                </Typography>
                <br />
                <Typography align="center">
                  <Link component={RouterLink} to="/Treatment">
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