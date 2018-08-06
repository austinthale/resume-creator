new Vue({
    el: '#app',
    data:{
        info: {
            person_info: {}
        },
        educations: [],
        employments: [],
        volunteers:[],
        isEditing: false
    },
    methods: {
        editButton: function() {
            this.isEditing = true;
        },
        saveButton: function() {
            axios
                .post('/resumejson', this.info.data,) //the data to post
                .then(r => console.log('r: ', JSON.stringify(r, null, 2)));
            this.isEditing = false;
        },
        addEduNote: function (idx) {
            console.log(this.info.educations);
            this.info.educations[idx].notes.push('');
        },
        deleteEduNote: function (i, j) {
            console.log(j);
            console.log(this.info.educations);
            this.info.educations[i].notes.splice(j, 1);
        },
        addJobNote: function (idx) {
            console.log(this.info.employments);
            this.info.employments[idx].notes.push('');
        },
        deleteJobNote: function (i, j) {
            console.log(j);
            console.log(this.info.employments);
            this.info.employments[i].notes.splice(j, 1);
        },
        addVolNote: function (idx) {
            console.log(this.info.volunteers);
            this.info.volunteers[idx].notes.push('');
        },
        deleteVolNote: function (i, j) {
            console.log(j);
            console.log(this.info.volunteers);
            this.info.volunteers[i].notes.splice(j, 1);
        },
        addEducation: function()  {
            var newEdu = {
                school: "",
                date_attended: "",
                notes: []
            };
            this.info.educations.push(newEdu);
        },
        addJob: function()  {
            var newJob = {
                company: "",
                position: "",
                date_attended: "",
                notes: []
            };
            this.info.employments.push(newJob);
        },
        addVol: function()  {
            var newVol = {
                company: "",
                position: "",
                date_attended: "",
                notes: []
            };
            this.info.volunteers.push(newVol);
        },
        deleteEdu: function (i) {
            this.info.educations.splice(i,1);
        },
        deleteJob: function (i) {
            this.info.employments.splice(i,1);
        },
        deleteVol: function (i) {
            this.info.volunteers.splice(i,1);
        },
    },
    mounted () {
        axios
            .get('/resumejson')
            .then(response => (this.info = response.data))
    }
});