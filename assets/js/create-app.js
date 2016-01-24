var CreateApp = angular.module('CreateApp', ['ui.materialize']);

CreateApp.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

CreateApp.controller('CreateEventCtrl', function($scope, $http) {
  $scope.newEvent = function() {
    return {
      name: "",
      event_type: "tall",
      date: "24 January, 2016",
      startTime: "1:00 PM",
      endTime: "2:00 PM",
      instructions: "",
      questions: [
        { type: "text", title: "Name", required: true, key: "asdf", choices: [] },
        { type: "text", title: "Email", required: true, key: "fdsa", choices: [] },
      ]
    }
  };

  $scope.eventData = $scope.newEvent();

  $scope.step = {
    title: "What Kind of Event?",
    key: 'event_type',
  };
  $scope.setType = function(val) {
    if ($scope.eventData.name != "") {
      $(".datepicker_mod").pickadate({clear:false});
      $(".timepicker_mod").pickatime();
      $scope.eventData['event_type'] = val;
      $scope.step = {
        title: "When is the Event?",
        key: "event_when_single",
      };
    }
  };
  $scope.selectSingleTime = function() {
    $scope.step = {
      title: "Need Instructions for Attendees?",
      key: "event_instructions",
    };
  };
  $scope.setInstructions = function(choice) {
    if (choice == 'blank') {
      $scope.eventData.instructions = "";
    }
    $scope.step = {
      title: "What Info Do You Need?",
      key: "event_questions",
    };
  };
  $scope.setQuestions = function() {
    $scope.step = {
      title: "Review Your Event!",
      key: "event_review",
    };
  };
  $scope.newQuestion = function() {
    $scope.eventData.questions.push({
      type: "text",
      title: "New Question",
      required: false,
      key: (Math.random()*1000).toFixed(),
    })
  };
  $scope.questionUp = function(key) {
    var idx = _.findIndex($scope.eventData.questions, {key: key});
    if (idx > 0) {
      var val = _.pullAt($scope.eventData.questions, idx);
      $scope.eventData.questions.splice(idx-1,0,val[0]);
    }
  };
  $scope.questionDown = function(key) {
    var idx = _.findIndex($scope.eventData.questions, {key: key});
    if (idx + 1 < $scope.eventData.questions.length) {
      var val = _.pullAt($scope.eventData.questions, idx);
      $scope.eventData.questions.splice(idx+1,0,val[0]);
    }
  };
  $scope.questionDelete = function(key) {
    var idx = _.findIndex($scope.eventData.questions, {key: key});
    _.pullAt($scope.eventData.questions, idx);
  };

  $scope.saveEvent = function() {
    $http.post(
      "/create_event",
      $scope.eventData
    ).then(function(data) {
      $scope.eventData = $scope.newEvent();
      $scope.eventKey = data.data.key;
      $scope.step = {
        title: "Event Created!",
        key: "event_done",
      };
    })
  };
});
