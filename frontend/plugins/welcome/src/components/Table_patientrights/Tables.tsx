import React, { useState, useEffect } from 'react';
import {  makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';


import { DefaultApi } from '../../api/apis';
import { EntPatientrights } from '../../api/models/EntPatientrights';


import {
  Grid,
  Typography,
  Avatar,
  Button,
} from '@material-ui/core';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';

import ImageIcon from '@material-ui/icons/Image';
import WorkIcon from '@material-ui/icons/Work';
import BeachAccessIcon from '@material-ui/icons/BeachAccess';
import Divider from '@material-ui/core/Divider';


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



export default function ComponentsTable(sim: any) {
  const classes = useStyles();
  const api = new DefaultApi();
  const [Patientrightss, setPatientrightss] = useState<EntPatientrights[]>(Array);
  const [loading, setLoading] = useState(true);
  const [nc, setnc] = useState(true);
  const [no, setno] = useState<EntPatientrights>({});

  const getPatientrightss = async () => {
    const res = await api.listPatientrights({ limit: 100, offset: 0 });
    setLoading(false);
    setPatientrightss(res);
    
  };

  useEffect(() => {
    
    getPatientrightss();
    setnc(true);
  }, [loading]);



  const Show = async (i: any) => {
    
    setno(i);
    setnc(false);
    
  };
  var p = 0;

  for (var val of Patientrightss) {
    if (val.edges?.edgesOfPatientrightsPatientrecord?.id === sim.sim || sim.sim === 0) {
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
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">ชื่อ ประกัน</TableCell>
                <TableCell align="center">Medicalrecordstaff</TableCell>
                <TableCell align="center">ความสามารถสิทธิ์</TableCell>
                <TableCell align="center">ผู้ป่วย</TableCell>
                <TableCell align="center">ผู้ดูแล</TableCell>
                <TableCell align="center">วันที่สร้าง(ปี-เดือน-วัน)</TableCell>
                <TableCell align="center">Manage</TableCell>

              </TableRow>
            </TableHead>
            <TableBody>
              {Patientrightss === undefined
                ? null
                : Patientrightss.filter(i => i.edges?.edgesOfPatientrightsPatientrecord?.id === sim.sim || sim.sim === 0).map((item: any) => (
                  <TableRow>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfPatientrightsInsurance?.insurancecompany}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfPatientrightsMedicalrecordstaff?.name}</TableCell>

                    <TableCell align="center">{item.edges?.edgesOfPatientrightsAbilitypatientrights?.check}</TableCell>
                    <TableCell align="center">{item.edges?.edgesOfPatientrightsPatientrecord?.name}</TableCell>
                    <TableCell align="center">{item.responsible}</TableCell>
                    <TableCell align="center">{item.permissionDate}</TableCell>
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
                <TableCell align="center">ไม่พบสิทธ์</TableCell>
              </TableRow>

            </TableHead>

          </Table>
        </TableContainer>
      );
    }
  } else {
    return (
      
      <List className={classes.root}>
        <div className={classes.section1}>
        <Grid container alignItems="center">
          <Grid item xs>
            <Typography gutterBottom variant="h4">
              Toothbrush
            </Typography>
          </Grid>
          <Grid item>
            <Typography gutterBottom variant="h6">
              $4.50
            </Typography>
          </Grid>
        </Grid>
        <Typography color="textSecondary" variant="body2">
          Pinstriped cornflower blue cotton blouse takes you on a walk to the park or just down the
          hall.
        </Typography>
      </div>
        <Divider variant="inset" component="li" />
        <ListItem>
          <ListItemAvatar>
            <Avatar>
              <WorkIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary="Work" secondary="Jan 7, 2014" />
        </ListItem>
        <Divider variant="inset" component="li" />
        <ListItem>
          <ListItemAvatar>
            <Avatar>
              <BeachAccessIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary="Vacation" secondary="July 20, 2014" />
        </ListItem>


        <Divider variant="inset" component="li" />
        <ListItem>
        <div className={classes.section3}>
        <Button color="primary" onClick={() => {setnc(true)}}>กลับ</Button>
      </div>
        </ListItem>

      </List>
     
    );
  }


}

