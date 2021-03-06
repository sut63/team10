import React, { useState, useEffect } from 'react';
import {
  Content,
  Page,
  Header,
  pageTheme,
} from '@backstage/core';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { Cookies } from 'react-cookie/cjs';//cookie
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntPrename } from '../../api/models/EntPrename';
import { EntEducationlevel } from '../../api/models/EntEducationlevel';
import { EntDepartment} from '../../api/models/EntDepartment';
import { EntOfficeroom } from '../../api/models/EntOfficeroom';
import { EntUser } from '../../api/models/EntUser';
import { Grid,TextField, Avatar } from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const cookies = new Cookies();
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
  
  telephonenumber?: string;
}

export default function CreateDoctorinfo() {
  const classes = useStyles();
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [Users, setUsers] = React.useState<Partial<EntUser>>();

  const [prenames, setPrenames] = useState<EntPrename[]>([]);
  const [nameError, setnameError] = React.useState('');
  const [lastnameError, setlastnameError] = React.useState('');
  const [telphonenumberError, settelphonenumberError] = React.useState('');
  const [doctornumberError, setdoctornumberError] = React.useState('');
  const [educationlevels, setEducationlevels] = useState<EntEducationlevel[]>([]);
  const [departments, setDepartments] = useState<EntDepartment[]>([]);
  const [officerooms, setOfficerooms] = useState<EntOfficeroom[]>([]);
  const [Doctorinfo, setDoctorinfo] = React.useState<Partial<Doctorinfo_Type>>({});
  const [loading, setLoading] = useState(true);

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
    const getPrenames = async () => {
      const res = await api.listPrename({ limit: 5, offset: 0 });
      setLoading(false);
      setPrenames(res);
    };
    getPrenames();

    const getEducationlevels = async () => {
      const res = await api.listEducationlevel({ limit: 3, offset: 0 });
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
      const res = await api.listOfficeroom({ limit: 5, offset: 0 });
      setLoading(false);
      setOfficerooms(res);
    };
    getOfficerooms();

    const getImg = async () => {
      const res = await api.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
  };
  getImg();

  }, [loading]);

const handleChange = (

  event: React.ChangeEvent<{ name?: string; value: any }>,
) => {
  const name = event.target.name as keyof typeof CreateDoctorinfo;
  const { value } = event.target;
  const validateValue = value.toString()
  checkPattern(name, validateValue)
  setDoctorinfo({ ...Doctorinfo, [name]: value });
};

// ฟังก์ชั่นสำหรับ validate ชื่อต้น
const validatename = (val: string) => {
  return val.charCodeAt(0) >= 65 && val.charCodeAt(0) <= 90 ? true : false;
}

// ฟังก์ชั่นสำหรับ validate นามสกุล
const validatesurname = (val: string) => {
  return val.charCodeAt(0) >= 65 && val.charCodeAt(0) <= 90 ? true : false;
}

// ฟังก์ชั่นสำหรับ validate เลขโทรศัพท์
const validatetelphoneNumber = (val: string) => {
  return val.length == 10 ? true : false;
}

// ฟังก์ชั่นสำหรับ validate เลขใบอนุญาติ
const validatedoctorNumber = (val: string) => {
  return val.length == 11 ? true : false;
}

// สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
const checkPattern  = (id: string, value: string) => {
  switch(id) {
    case 'doctorname':
      validatename(value) ? setnameError('') : setnameError('ชื่อต้องขึ้นต้นด้วยอักษรพิมพ์ใหญ่');
      return;
    case 'doctorsurname':
      validatesurname(value) ? setlastnameError('') : setlastnameError('นามสกุลต้องขึ้นต้นด้วยอักษรพิมพ์ใหญ่');
      return;
    case 'telephonenumber':
      validatetelphoneNumber(value) ? settelphonenumberError('') : settelphonenumberError('หมายเลขโทรศัพท์ 10 หลัก')
      return;
    case 'licensenumber':
      validatedoctorNumber(value) ? setdoctornumberError('') : setdoctornumberError('หมายเลขใบประกอบวิชาชีพ 11 หลัก')
      return;
    default:
      return;
  }
}

const alertMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}

const checkCaseSaveError = (field: string) => {
  switch(field) {
    case 'doctorname':
      alertMessage("error","ชื่อต้องขึ้นต้นด้วยอักษรพิมพ์ใหญ่");
      return;
    case 'doctorsurname':
      alertMessage("error","นามสกุลต้องขึ้นต้นด้วยอักษรพิมพ์ใหญ่");
      return;
    case 'telephonenumber':
      alertMessage("error","หมายเลขโทรศัพท์ 10 หลัก");
      return;
    case 'licensenumber':
      alertMessage("error","หมายเลขใบประกอบวิชาชีพ 11 หลัก");
      return;
    default:
      alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
      return;
  }
}

const Create_Doctorinfo = async () => {
  const apiUrl = 'http://localhost:8080/api/v1/doctorinfos';
  const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(Doctorinfo),
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
        checkCaseSaveError(data.error.Name)
      }
    });
};

  {/*const Create_Doctorinfo = async () => {
   
   if ((Doctorinfo.department != null) && (Doctorinfo.doctorname != '')
     && (Doctorinfo.doctorsurname != '') && (Doctorinfo.educationlevel != null)
     && (Doctorinfo.licensenumber != '')&& (Doctorinfo.officeroom != null)
     && (Doctorinfo.prename != null)&& (Doctorinfo.telephonenumber != '')) {
   const res: any = await api.createDoctorinfo({ 
      doctorinfo:Doctorinfo
    
    
    });
    console.log(Doctorinfo);
    
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
  }; */}

  const profile = { givenName: '' };
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`DOCTOR INFORMATION DEPARTMENT`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
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
                    style={{ width: 250 }}
                >
                {prenames.map((item: any) => (
                  <MenuItem value={item.id}>{item.prefix}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>

                
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="doctorname"
                                    label="ชื่อ"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.doctorname}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
                        
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="doctorsurname"
                                    label="นามสกุล"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.doctorsurname}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>

                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="telephonenumber"
                                    label="เบอร์ติดต่อ"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.telephonenumber}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
       
                        <div className={classes.root}>
                            <form noValidate autoComplete="off">
                            <FormControl variant="filled" className={classes.formControl}>
                                <TextField
                                    name="licensenumber"
                                    label="หมายเลขใบประกอบวิชาชีพ"
                                    variant="outlined"
                                    type="string"
                                    size="medium"
                                        
                                    value={Doctorinfo.licensenumber}
                                    onChange={handleChange}
                                />
                            </FormControl>
                            </form>
                        </div><br/>
                </Typography>

                <Typography variant="h6" gutterBottom  align="center">
                    เลือกระดับการศึกษา
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="educationlevels"
                    id="educationlevels"
                    name="educationlevel"
                    value={Doctorinfo.educationlevel}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {educationlevels.map((item: EntEducationlevel) => (
                  <MenuItem value={item.id}>{item.level}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>

                <Typography variant="h6" gutterBottom  align="center">
                    เลือกเเผนกของคุณ
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="departments"
                    id="departments"
                    name="department"
                    value={Doctorinfo.department}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {departments.map((item: any) => (
                    <MenuItem value={item.id}>{item.department}</MenuItem>
                ))}
                </Select>
                </Typography>
                </Typography><br/>

                <Typography variant="h6" gutterBottom  align="center">
                    เลือกห้องพักของคุณ 
                <Typography variant="body1" gutterBottom> 
                <Select
                    labelId="officerooms"
                    id="officerooms"
                    name="officeroom"
                    value={Doctorinfo.officeroom}
                    onChange={handleChange}
                    style={{ width: 250 }}
                >
                {officerooms.map((item: any) => (
                  <MenuItem value={item.id}>{item.roomnumber}</MenuItem>
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
                            <Alert severity="error" style={{ marginTop: 40 }}>
                            Failed to save!!! Please check again.
                            </Alert>
                        )}
                    </div>
                    ) : null}
                </Grid> <br/>

                <div className={classes.margin}>
                <Typography variant="h6" gutterBottom  align="center">
                <Button
                    onClick={() => {
                    Create_Doctorinfo();
                    }}
                    variant="contained"
                    color="primary"
                >
                    Submit
                </Button>
                </Typography> <br/>
                </div>

                <div className={classes.margin}>
                <Typography variant="h6" gutterBottom  align="center">
                <Button
                style={{ marginLeft: 40 }}
                component={RouterLink}
                to="/Table_Doctors"
                variant="contained"
                color="secondary"
              >
                SHOW
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
