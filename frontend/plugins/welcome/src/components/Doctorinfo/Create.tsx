import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Page,
  BootErrorPageProps,
  Header,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { EntPrename } from '../../api/models/EntPrename';
import { EntEducationlevel } from '../../api/models/EntEducationlevel';
import { EntDepartment} from '../../api/models/EntDepartment';
import { EntOfficeroom } from '../../api/models/EntOfficeroom';
import { Grid, Paper, TextField } from '@material-ui/core';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    formControl: {
        width: 250,
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

interface Doctorinfo_Type {
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  department?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  doctorname?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  doctorsurname?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  educationlevel?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  licensenumber?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  officeroom?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  prename?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  registrar?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersDoctorinfo
   */
  telephonenumber?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersDoctorinfo
   */
  user?: number;
}

export default function CreateDoctorinfo() {
  const classes = useStyles();
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [prenames, setPrenames] = useState<EntPrename[]>([]);
  const [educationlevels, setEducationlevels] = useState<EntEducationlevel[]>([]);
  const [departments, setDepartments] = useState<EntDepartment[]>([]);
  const [officerooms, setOfficerooms] = useState<EntOfficeroom[]>([]);

  const [Doctorinfo, setDoctorinfo] =            React.useState<Partial<Doctorinfo_Type>>({});

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getPrenames = async () => {
      const res = await api.listPrename({ limit: 3, offset: 0 });
      setLoading(false);
      setPrenames(res);
    };
    getPrenames();

    const getEducationlevels = async () => {
      const res = await api.listEducationlevel({ limit: 2, offset: 0 });
      setLoading(false);
      setEducationlevels(res);
      console.log(res);
    };
    getEducationlevels();

    const getDepartments = async () => {
      const res = await api.listDepartment({ limit: 3, offset: 0 });
      setLoading(false);
      setDepartments(res);
    };
    getDepartments();

    const getOfficerooms = async () => {
      const res = await api.listOfficeroom({ limit: 3, offset: 0 });
      setLoading(false);
      setOfficerooms(res);
    };
    getOfficerooms();

  }, [loading]);

const handleChange = (

  event: React.ChangeEvent<{ name?: string; value: unknown }>,
) => {
  const name = event.target.name as keyof typeof CreateDoctorinfo;
  const { value } = event.target;
  setDoctorinfo({ ...Doctorinfo, [name]: value });
};

  const CreateDoctorinfo = async () => {
    const res: any = await api.createDoctorinfo({ 
      doctorinfo:Doctorinfo
    
    
    });
    console.log(Doctorinfo);
    
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    setTimeout(() => {
      setStatus(false);
    }, 3000);
  };
  const profile = { givenName: '' };
  return (
    <Page theme={pageTheme.home}>
      <Header
      title={`${profile.givenName || 'DOCTOR INFORMATION DEPARTMENT'}`}
      subtitle=""
     ></Header>
      <Content>
        <Grid container spacing = {5} >
          <Grid container item xs = {12} sm = {12}  >
              <Grid item xs = {12}>
                <Typography align ="center">
                    <Typography align = "center" variant = "h3">
                      <br/>=====  Create Doctorinformation  ===== <br/> <br/>
                    </Typography> 
                    <Typography variant="h6" gutterBottom  align="center">
                     เลือกคำนำหน้าชื่อ
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="prenames"
                    id="prenames"
                    name="prename"
                    value={Doctorinfo.prename}
                    onChange={handleChange}
                    style={{ width: 500 }}
                >
                {patientrecords.map((item: any) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>


                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="hight"
                                    label="Hight(cm)"
                                    variant="outlined"
                                    type="number"
                                    size="medium"
                                        
                                    value={Historytaking.hight}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
      
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="weight"
                                    label="Weight(kg)"
                                    variant="outlined"
                                    type="number"
                                    size="medium"
                                    
                                    value={Historytaking.weight}
                                    onChange={handleChange}
                                    />
                            </FormControl>
                            </form>
                        </div><br/>
                        
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                            <TextField
                                    name="temp"
                                    label="Temperature(Celcius)"
                                    variant="outlined"
                                    type="number"
                                    size="medium"
                                    
                                    value={Historytaking.temp}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>

                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                            <TextField
                                name="pulse"
                                label="Pulse(Times/minute)"
                                variant="outlined"
                                type="number"
                                size="medium"
                                
                                value={Historytaking.pulse}
                                onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>

                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                            <TextField
                                name="respiration"
                                label="Respiration(Times/minute)"
                                variant="outlined"
                                type="number"
                                size="medium"
                                
                                value={Historytaking.respiration}
                                onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                            <TextField
                                name="bp"
                                label="Blood pressure(mm/Hg)"
                                variant="outlined"
                                type="string"
                                size="medium"
                                
                                value={Historytaking.bp}
                                onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
       
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                            <TextField
                                name="oxygen"
                                label="Oxygen(%)"
                                variant="outlined"
                                type="string"
                                size="medium"
                                
                                value={Historytaking.oxygen}
                                onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                            <TextField
                                name="symptom"
                                label="Symptomseverity"
                                variant="outlined"
                                type="string"
                                size="medium"
                                
                                value={Historytaking.symptom}
                                onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
                </Typography>

                <Typography variant="h6" gutterBottom  align="center">
                    Nurses ID : 
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="nurses"
                    id="nurses"
                    name="nurse"
                    value={Historytaking.nurse}
                    onChange={handleChange}
                    style={{ width: 500 }}
                >
                {nurses.map((item: EntNurse) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>
                <Typography variant="h6" gutterBottom  align="center">
                    Symptomseverity ID : 
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="symptomseveritys"
                    id="symptomseveritys"
                    name="symptomseverity"
                    value={Historytaking.symptomseverity}
                    onChange={handleChange}
                    style={{ width: 500 }}
                >
                {symptomseveritys.map((item: any) => (
                  <MenuItem value={item.id}>{item.symptomseverity}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>
                <Typography variant="h6" gutterBottom  align="center">
                    Department ID : 
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="departments"
                    id="departments"
                    name="department"
                    value={Historytaking.department}
                    onChange={handleChange}
                    style={{ width: 500 }}
                >
                {departments.map((item: any) => (
                    <MenuItem value={item.id}>{item.department}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>
                <Typography variant="h6" gutterBottom  align="center">
                    Patientrecord ID : 
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="patientrecords"
                    id="patientrecords"
                    name="patientrecord"
                    value={Historytaking.patientrecord}
                    onChange={handleChange}
                    style={{ width: 500 }}
                >
                {patientrecords.map((item: any) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>
                <Grid container justify="center">
                    {status ? (
                    <div>
                        {alert ? (
                        <Alert severity="success">
                            success!
                        </Alert>
                        ) : (
                            <Alert severity="warning" style={{ marginTop: 40 }}>
                            This is a warning alert — check it out!
                            </Alert>
                        )}
                    </div>
                    ) : null}
                </Grid> <br/>

                <div className={classes.margin}>
                <Typography variant="h6" gutterBottom  align="center">
                <Button
                    onClick={() => {
                    CreateHistorytaking();
                    }}
                    variant="contained"
                    color="primary"
                >
                    Submit
                </Button>
                <Button
                    style={{ marginLeft: 40 }}
                    component={RouterLink}
                    to="/bill"
                    variant="contained"
                >
                    Back
                </Button>
                </Typography> <br/>
                </div>
                </Grid>
            </Grid>
        </Grid>     
      </Content>
    </Page>
  );
};
