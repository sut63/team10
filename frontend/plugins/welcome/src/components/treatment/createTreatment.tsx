import React, { useEffect ,FC } from 'react';
import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {Grid,MenuItem,Button,TextField,FormControl,Select, Typography} from '@material-ui/core';

import { Alert } from '@material-ui/lab';
import Paper from '@material-ui/core/Paper';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';


import { DefaultApi } from'../../api/apis';
import { EntTypetreatment } from '../../api/models/EntTypetreatment';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { EntTreatment } from '../../api/models/EntTreatment';

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
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    const [patientrecords, setPatientrecords] = React.useState<EntPatientrecord[]>([]);
    const [treatments, setTreatments] = React.useState<EntTreatment[]>([]);

    const [treatmentes, settreatment] = React.useState(String);
    const [datetreat, setDatetreat] = React.useState(String);
    const [doctorid, setdoctorId] = React.useState(Number);
    const [patientrecordid, setpatientrecordId] = React.useState(Number);
    const [typetreatmentid, settypetreatmentId] = React.useState(Number);

    useEffect(() => {
        const getDocdor = async () => {
            const res = await http.listDoctor({ limit: 100, offset: 0 });
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
        getPatientrecord();
        getTypetreatment();
        getTreatment();
        getDocdor();
    }, [loading]);

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
          const tm = {
            treatment: treatmentes,
            datetreat: datetreat,
            typetreatment: typetreatmentid,
            doctor: doctorid,
            patientrecord: patientrecordid,
          };
          console.log(tm);
          const res: any = await http.createTreatment({ treatment : tm });
          setStatus(true);
          if (res.id != '') {
            setAlert(true);
          } else {
            setAlert(false);
          }
          setTimeout(() => {
            setStatus(false);
          }, 1000);
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
             มีข้อผิดพลาด โปรดลองอีกครั้ง
           </Alert>
         )}
     </div>
    ):null}


    <Page theme={pageTheme.home}>
        <Header title={`Treatment Department`}></Header>
      <Content>
        <Grid container spacing = {3} >
          <Grid container item xs = {12} sm = {12}  >
              <Grid item xs = {4}>
                  <Paper >
                <Typography align ="center">
                    <Typography align = "center" variant = "h3">
                      <br/>----  Create Traetment  ----
                    </Typography>
                  <FormControl className={classes.formControl}>
                        <Typography align = "center" variant = "h6">
                        <br/>รูปแบบการรักษา
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
                <Typography align = "center" variant = "h6">
                     <br/>รายละเอียดการรักษา<br/>
                <TextField 
                    className={classes.formControl}
                    value={treatmentes}
                    onChange={TreatmenthandleChange}/>
                </Typography>
                <FormControl className={classes.formControl}>
                        <Typography align = "center"variant = "h6">
                        <br/>แพทย์
                        </Typography>
                        <Select
                        name="doctor"
                        value={doctorid}
                        onChange={DoctorhandleChange}
                        >
                        {doctors.map(item => {
                        return (
                        <MenuItem value={item.id}>
                         {item.edges?.doctorinfo?.doctorname} {item.edges?.doctorinfo?.doctorsurname}
                        </MenuItem>
                        );
                        })}
                        </Select>
                </FormControl>
                <br/>
                <FormControl className={classes.formControl}>
                        <Typography align = "center"variant = "h6">
                        <br/>ผู้ป่วย
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
                <br/>
              
                <Typography align ="center">
                <br/>
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
            </Typography>
            <br/>
            </Paper>
            <Paper>
            </Paper>
              </Grid>
              <Grid item xs = {8}>
                  <Paper>
                  <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center" >เลขที่การรักษา</TableCell>
         <TableCell align="center">แพทย์</TableCell>
         <TableCell align="center">ผู้เข้ารับการรักษา</TableCell>
         <TableCell align="center">รูปแบบการรักษา</TableCell>
         <TableCell align="center">วันเวลาที่รักษา</TableCell>
         </TableRow>
         </TableHead>
         <TableBody>
         {treatments.map(item => (doctors.filter(t => t.id === item.edges?.doctor?.id).map(item2 => (
                <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item2.edges?.doctorinfo?.doctorname} {item2.edges?.doctorinfo?.doctorsurname}</TableCell>
                    <TableCell align="center">{item.edges?.patientrecord?.name}</TableCell>
                    <TableCell align="center">{item.edges?.typetreatment?.typetreatment}</TableCell> 
                    <TableCell align="center">{item.datetreat}</TableCell> 
                      </TableRow>
         ))))}
         </TableBody>
         </Table>
        </TableContainer>
                  </Paper>
                  </Grid>
          </Grid>
        </Grid>
      </Content>
    </Page>
    </div>
  );
};

export default createTreatment;