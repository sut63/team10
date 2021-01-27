import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Grid, Typography, Button, Link} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis';
import { EntHistorytaking } from '../../api';


const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 200,
  },
  root: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,

  },
  div: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  section1: {
    margin: theme.spacing(3, 2),
  },
  section3: {
    margin: theme.spacing(3, 1, 1),
  },
}));

const refreshPage = () => {
  window.location.reload();
}

export default function ComponentsTable(sim: any) {
  const classes = useStyles();
  const api = new DefaultApi();
  const [Historytakings, setHistorytakings] = useState<EntHistorytaking[]>(Array);
  const [loading, setLoading] = useState(true);
  const [nc, setnc] = useState(true);
  const [no, setno] = useState<EntHistorytaking>({});

  const getHistorytakings = async () => {
    const res = await api.listHistorytaking({ limit: 100, offset: 0 });
    setLoading(false);
    setHistorytakings(res);

  };

  useEffect(() => {

    getHistorytakings();
    setnc(true);
  }, [loading]);



  const Show = async (i: any) => {

    setno(i);
    setnc(false);

  };
  var p = 0;

  for (var val of Historytakings) {
    if (val.edges?.edgesOfPatientrecord?.name === sim.sim || sim.sim === undefined || sim.sim === null) {
      p = p + 1
    }

  }
  if (nc) {

    if (p !== 0) {
      return (

        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">ID</TableCell>
                <TableCell align="center">Hight</TableCell>
                <TableCell align="center">Weight</TableCell>
                <TableCell align="center">T</TableCell>
                <TableCell align="center">P</TableCell>
                <TableCell align="center">R</TableCell>
                <TableCell align="center">BP</TableCell>
                <TableCell align="center">O2</TableCell>
                <TableCell align="center">Symptom</TableCell>
                <TableCell align="center">Nurse</TableCell>
                <TableCell align="center">Symptomseverity</TableCell>
                <TableCell align="center">Department</TableCell>
                <TableCell align="center">Patient</TableCell>
                <TableCell align="center">Datetime</TableCell>
                <TableCell align="center">Manage</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {Historytakings === undefined
                ? null
                : Historytakings.filter(i => i.edges?.edgesOfPatientrecord?.name === sim.sim || sim.sim === undefined || sim.sim === null).map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item.hight}</TableCell>
                    <TableCell align="center">{item.weight}</TableCell>
                    <TableCell align="center">{item.temp}</TableCell>
                    <TableCell align="center">{item.pulse}</TableCell>
                    <TableCell align="center">{item.respiration}</TableCell>
                    <TableCell align="center">{item.bp}</TableCell>
                    <TableCell align="center">{item.oxygen}</TableCell>
                    <TableCell align="center">{item.symptom}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfNurse?.name}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfSymptomseverity?.symptomseverity}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfDepartment?.department}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfPatientrecord?.name}</TableCell>
                    <TableCell align="center">{item.datetime}</TableCell>
                    <TableCell align="center">
                      <Button
                        onClick={() => {
                          Show(item);
                        }}
                        style={{ marginLeft: 10 }}
                        variant="contained"
                        color="secondary"
                      >
                        ดูข้อมูล
               </Button>
                    </TableCell>
                  </TableRow>
                ))}
            </TableBody>
          </Table>
        </TableContainer>
      );
    } else {
      return (
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">ไม่พบบันทึกการซักประวัติผู้ป่วยนอก</TableCell>
              </TableRow>

            </TableHead>

          </Table>
        </TableContainer>
      );
    }
  } else {
    return (
      <TableContainer component={Paper}>
        <Typography align="center" variant="h1">
          Outpatient history taking records: 
        </Typography>
          <Grid item xs={12}>
              <Typography align="center" variant="h5">
                <br/>
                Patientname:&emsp;{no.edges?.edgesOfPatientrecord?.name}&emsp;by&emsp;{no.edges?.edgesOfNurse?.name}<br/>
                Department:&emsp;{no.edges?.edgesOfDepartment?.department}<br/>
                Symptomseverity:&emsp;{no.edges?.edgesOfSymptomseverity?.symptomseverity}<br/>
              </Typography>
          </Grid><br/>
          <Grid item xs={12}>
            <Typography align="center" variant="h6">
              Hight:&emsp;{no.hight}&emsp;<br/>
              Weight:&emsp;{no.weight}&emsp;<br/>
              Temperature:&emsp;{no.temp}&emsp;<br/>
              Pulse:&emsp;{no.pulse}&emsp;<br/>
              Respiration:&emsp;{no.respiration}&emsp;<br/>
              Blood pressure:&emsp;{no.bp}&emsp;<br/>
              %Oxygen:&emsp;{no.oxygen}<br/>
              Sympton:&emsp;{no.symptom}<br/>
              Datetime:&emsp;{no.datetime}<br/><br/>
            </Typography>
            <Typography align="center" variant="h6">
            <Link component={RouterLink} to="/tableHistorytaking">
            <Button 
              variant="contained"
              onClick={() => {
                refreshPage();
              }}>
              BACK
            </Button>
            </Link>
            </Typography><br/><br/>
          </Grid>
      </TableContainer>
    );
  }
}

