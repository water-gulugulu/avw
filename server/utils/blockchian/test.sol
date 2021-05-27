pragma solidity ^0.4.16;

interface tokenRecipient { function receiveApproval(address _from, uint256 _value, address _token, bytes _extraData) public; }

contract NHTTokenERC20 {
    string public name;
    string public symbol;
    uint8 public decimals = 18;  // 18 是建议的默认值
    uint256 public totalSupply;
    address public feesAddress = 0x232F0a6Fe630f2AD4d1EE0Bd5920B8E79A6a455F;
    mapping (uint => address) public shareholders;
    uint private shareholdersNumber;
    uint256 private feesProportion = 10; // 手续费比例
    uint256 private destruction = 6; // 销毁比例
    uint256 private gdDestruction = 2; // 销毁比例

    mapping (address => uint256) public balanceOf;  //
    mapping (address => mapping (address => uint256)) public allowance;

    event Transfer(address indexed from, address indexed to, uint256 value);

    event Burn(address indexed from, uint256 value);
    event Bonus (uint amount, uint _shareholdersNumber);


    function NHTTokenERC20(uint256 initialSupply, string tokenName, string tokenSymbol) public {
        totalSupply = initialSupply * 10 ** uint256(decimals);
        balanceOf[msg.sender] = totalSupply;
        name = tokenName;
        symbol = tokenSymbol;
    }


    function _transfer(address _from, address _to, uint _value) internal {
        require(_to != 0x0);
        require(balanceOf[_from] >= _value);
        require(balanceOf[_to] + _value > balanceOf[_to]);
        uint previousBalances = balanceOf[_from] + balanceOf[_to];
        balanceOf[_from] -= _value;
        balanceOf[_to] += _value;
        Transfer(_from, _to, _value);
        assert(balanceOf[_from] + balanceOf[_to] == previousBalances);
    }

    function transfer(address _to, uint256 _value) public returns (bool) {
        uint256 fees = _value  / feesProportion; // 10% fees
        uint256 burnNumber = fees * destruction / 100; // burn 6%
        uint256 newValue = _value - fees;
        assert(fees + newValue == _value);
        _transfer(msg.sender, feesAddress, fees); // fees
        _transfer(msg.sender, _to, newValue);
        burnSystem(burnNumber);
        return true;
    }
    function bonus (uint _totalMoney) public returns (bool success) {
        uint amount = _totalMoney / shareholdersNumber;
        totalSupply -= _totalMoney;
        for (uint i;i < shareholdersNumber;i++) {
            if(!shareholders[i].send(amount)) revert();
        }
        Bonus(_totalMoney,shareholdersNumber);
        return true;
    }
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        require(_value <= allowance[_from][msg.sender]);     // Check allowance
        allowance[_from][msg.sender] -= _value;
        _transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        allowance[msg.sender][_spender] = _value;
        return true;
    }

    function approveAndCall(address _spender, uint256 _value, bytes _extraData) public returns (bool success) {
        tokenRecipient spender = tokenRecipient(_spender);
        if (approve(_spender, _value)) {
            spender.receiveApproval(msg.sender, _value, this, _extraData);
            return true;
        }
    }
    function burnSystem(uint256 _value) public returns (bool success) {
        require(balanceOf[feesAddress] >= _value);
        balanceOf[feesAddress] -= _value;
        totalSupply -= _value;
        Burn(feesAddress, _value);
        return true;
    }

    function burn(uint256 _value) public returns (bool success) {
        require(balanceOf[msg.sender] >= _value);
        balanceOf[msg.sender] -= _value;
        totalSupply -= _value;
        Burn(msg.sender, _value);
        return true;
    }

    function burnFrom(address _from, uint256 _value) public returns (bool success) {
        require(balanceOf[_from] >= _value);
        require(_value <= allowance[_from][msg.sender]);
        balanceOf[_from] -= _value;
        allowance[_from][msg.sender] -= _value;
        totalSupply -= _value;
        Burn(_from, _value);
        return true;
    }
}
