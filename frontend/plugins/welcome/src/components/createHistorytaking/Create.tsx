import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Page,
  Header,
  pageTheme,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { EntNurse } from '../../api/models/EntNurse';
import { EntSymptomseverity } from '../../api/models/EntSymptomseverity';
import { EntDepartment } from '../../api/models/EntDepartment';
import { EntPatientrecord } from '../../api/models/EntPatientrecord';
import { Grid, TextField, Avatar } from '@material-ui/core';
import { Cookies } from 'react-cookie/cjs';//cookie
import { EntUser } from '../../api/models/EntUser';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
const NURID = cookies.get('Nur');
const Name = cookies.get('Name');
const Img = cookies.get('Img');

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

interface Historytaking_Type {
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  bp?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  datetime?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersHistorytaking
   */
  department?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  hight?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersHistorytaking
   */
  nurse?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  oxygen?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersHistorytaking
   */
  patientrecord?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  pulse?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  respiration?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  symptom?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersHistorytaking
   */
  symptomseverity?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  temp?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersHistorytaking
   */
  weight?: string;
}

export default function CreateHistorytaking() {
  const classes = useStyles();
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [nurses, setNurses] = React.useState<Partial<EntNurse>>();
  const [symptomseveritys, setSymptomseveritys] = useState<EntSymptomseverity[]>([]);
  const [departments, setDepartments] = useState<EntDepartment[]>([]);
  const [patientrecords, setPatientrecords] = useState<EntPatientrecord[]>([]);

  const [Historytaking, setHistorytaking] = React.useState<Partial<Historytaking_Type>>({});

  const [loading, setLoading] = useState(true);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  useEffect(() => {

    const getChangeOfUser = async () => {

      const name = "nurse";
      const value = parseInt(NURID, 10);
      setHistorytaking({ ...Historytaking, [name]: value });
    };
    getChangeOfUser();

    const getSymptomseveritys = async () => {
      const res = await api.listSymptomseverity({ limit: 3, offset: 0 });
      setLoading(false);
      setSymptomseveritys(res);
    };
    getSymptomseveritys();

    const getNurses = async () => {
      const res = await api.getNurse({ id: Number(NURID) });
      setLoading(false);
      setNurses(res);
      console.log(res);
    };
    getNurses();

    const getDepartments = async () => {
      const res = await api.listDepartment({ limit: 3, offset: 0 });
      setLoading(false);
      setDepartments(res);
    };
    getDepartments();

    const getImg = async () => {
      const res = await api.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };
    getImg();

    const getPatientrecords = async () => {
      const res = await api.listPatientrecord({ limit: 3, offset: 0 });
      setLoading(false);
      setPatientrecords(res);
    };
    getPatientrecords();


  }, [loading]);

  const handleChange = (

    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof CreateHistorytaking;
    const { value } = event.target;
    setHistorytaking({ ...Historytaking, [name]: value });
  };

  const CreateHistorytaking = async () => {

    if ((Historytaking.bp != '') && (Historytaking.datetime != '')
      && (Historytaking.hight != '') && (Historytaking.oxygen != '')
      && (Historytaking.pulse != '') && (Historytaking.symptom != '')
      && (Historytaking.temp != '') && (Historytaking.weight != '')
      && (Historytaking.bp != '') && (Historytaking.symptomseverity != null)
      && (Historytaking.patientrecord != null) && (Historytaking.nurse != null)
      && (Historytaking.department != null)) {

      const res: any = await api.createHistorytaking({
        historytaking: Historytaking
      });
      console.log(Historytaking);

      setStatus(true);
      if (res.id != '') {
        setAlert(true);
      }
    }
    else {
      setStatus(true);
      setAlert(false);
      setTimeout(() => {
        setStatus(false);
      }, 3000);
    }
  };
  const profile = { givenName: '' };
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`HISTORYTAKING DEPARTMENT`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <Grid container spacing={5} >
          <Grid container item xs={12} sm={12}  >
            <Grid item xs={12}>
              <Typography align="center">
                <Typography align="center" variant="h3">
                  <br />=====  Create Historytaking  ===== <br /> <br />
                </Typography>
                <div className={classes.root}>
                  <form noValidate autoComplete="off">
                    <FormControl variant="filled" className={classes.formControl}>
                      <TextField
                        name="hight"
                        label="Hight(cm)"
                        variant="outlined"
                        type="string"
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
                        type="string"
                        size="medium"

                        value={Historytaking.weight}
                        onChange={handleChange}
                      />
                    </FormControl>
                  </form>
                </div><br />

                <div className={classes.root}>
                  <form noValidate autoComplete="off">
                    <FormControl variant="filled" className={classes.formControl}>
                      <TextField
                        name="temp"
                        label="Temperature(Celcius)"
                        variant="outlined"
                        type="string"
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
                        type="string"
                        size="medium"

                        value={Historytaking.pulse}
                        onChange={handleChange}
                      />
                    </FormControl>
                  </form>
                </div><br />

                <div className={classes.root}>
                  <form noValidate autoComplete="off">
                    <FormControl variant="filled" className={classes.formControl}>
                      <TextField
                        name="respiration"
                        label="Respiration(Times/minute)"
                        variant="outlined"
                        type="string"
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
                </div><br />

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
                </div><br />
              </Typography>

              <Typography variant="h6" gutterBottom align="center">
                Nurse : {nurses?.name}
                <Typography variant="body1" gutterBottom>

                </Typography>
              </Typography><br />
              <Typography variant="h6" gutterBottom align="center">
                Symptomseverity :
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
              </Typography><br />
              <Typography variant="h6" gutterBottom align="center">
                Department :
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
              </Typography><br />
              <Typography variant="h6" gutterBottom align="center">
                Patientrecord :
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
              </Typography><br />
              <Grid container justify="center">
                {status ? (
                  <div>
                    {alert ? (
                      <Alert severity="success">
                        success!
                      </Alert>
                    ) : (
                        <Alert severity="error" style={{ marginTop: 40 }}>
                          This is a error alert — check it out!
                        </Alert>
                      )}
                  </div>
                ) : null}
              </Grid> <br />

              <div className={classes.margin}>
                <Typography variant="h6" gutterBottom align="center">
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
                    to="/"
                    variant="contained"
                  >
                    Back
                </Button>
                  <Button
                    style={{ marginLeft: 40 }}
                    component={RouterLink}
                    to="/Historytaking"
                    variant="contained"
                    color="secondary"
                  >
                    SHOW
                </Button>
                </Typography> <br />
              </div>
            </Grid>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};
