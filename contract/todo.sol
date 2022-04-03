pragma solidity >=0.7.0 <0.9.0;

contract Todo {
    address public owner;
    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isOwnwer() {
        require(owner == msg.sender);
        _; //执行方法内的逻辑
    }

    function add(string memory _content) public isOwnwer {
        tasks.push(Task(_content, false));
    }

    function get(uint256 _id) public view isOwnwer returns (Task memory) {
        return tasks[_id];
    }

    function list() public view isOwnwer returns (Task[] memory) {
        return tasks;
    }

    function update(uint256 _id, string memory _content) public isOwnwer {
        tasks[_id].content = _content;
    }

    function toggle(uint256 _id) public isOwnwer {
        tasks[_id].status = !tasks[_id].status;
    }

    function remove(uint256 _id) public isOwnwer {
        for (uint256 i = _id; i < tasks.length - 1; i++) {
            tasks[i] = tasks[i + 1];
        }
        tasks.pop();
    }
}
